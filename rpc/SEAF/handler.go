package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	"context"
	"fmt"
	"log"
	"time"
)

var (
	ausfClient ausf.Client
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context, data string) (resp string, err error) {
	// Judge data from UE: access-request or auth-response
	length := len(data)
	if length == 30 { // access-request
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive SUCI and snName from UE.")
		// send SUPI and SN_Name to AUSF, get response and update logs
		AV, err := ausfClient.Authenticate(context.Background(), data)
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and snName to AUSF.")
		if err != nil {
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
			return "", err
		}

		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive 5G_SE_AV and SUPI from AUSF.")
		var AUTN string
		randNum, AUTN, hxResStar, _ = ResolveAV(AV)

		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send rand_num and AUTN to UE.")

		// Send rand_num and AUTN to UE
		return randNum + AUTN, nil

	} else if length == 32 { // auth-response
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive res* from UE.")
		resStar := data
		hResStar := GenerateHResStar(randNum, resStar)

		// Judge
		if hResStar == hxResStar {
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "SEAF Authentication Passed.")

			_, err := ausfClient.Authenticate(context.Background(), resStar)
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send res* to AUSF.")
			if err != nil {
				log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
				return "", err
			}
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Authentication Passed! Send Authentication Response to UE.")
			return "Authentication passed successfully!", nil
		} else {
			msg := "SEAF Authentication failed: Unable to pass access-auth-check"
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + msg)
			return "", fmt.Errorf(msg)
		}
	} else {
		err = fmt.Errorf("SEAF Authentication failed: Unable to process unknown data")
		return "", err
	}

	return
}
