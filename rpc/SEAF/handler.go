package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	"context"
	"fmt"
	"os"
	"time"
)

var (
	logPath    string = "../../log/SEAF.log"
	ausfClient ausf.Client
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context, data string) (resp string, err error) {

	// Load log file
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Judge data from UE
	length := len(data)
	if length == 30 {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive SUCI and snName from UE.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive SUCI and snName from UE.")

		// send SUPI and SN_Name to AUSF, get response and update logs
		AV, err := ausfClient.Authenticate(context.Background(), data)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and snName to AUSF.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and snName to AUSF.")
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Failed to receive response from AUSF")
			err = fmt.Errorf(time.Now().Format("2006-01-02 15:04:05") + "  " + "Failed to receive response from AUSF")
			return "Failed to receive response from AUSF", err
		}

		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive 5G_SE_AV and SUPI from AUSF.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive 5G_SE_AV and SUPI from AUSF.")
		var AUTN string
		randNum, AUTN, hxResStar, _ = ResolveAV(AV)

		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send rand_num and AUTN to UE.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send rand_num and AUTN to UE.")

		// Send rand_num and AUTN to UE
		return randNum + AUTN, nil

	} else if length == 32 {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive res* from UE.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive res* from UE.")
		resStar := data
		hResStar := GenerateHResStar(randNum, resStar)

		// Judge
		if hResStar == hxResStar {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "SEAF Authentication Passed.")
			_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "SEAF Authentication Passed.")

			resp, err := ausfClient.Authenticate(context.Background(), resStar)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send res* to AUSF.")
			_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send res* to AUSF.")
			if err != nil {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + resp)
				_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + resp)
				return resp, fmt.Errorf(resp)
			}
		} else {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "SEAF Authentication Failed!")
			_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "SEAF Authentication Failed!")
			return "SEAF Authentication Failed!", fmt.Errorf("SEAF Authentication Failed!")
		}
	} else {
		msg := "Unknown data"
		err = fmt.Errorf(msg)
		return msg, err
	}

	return
}
