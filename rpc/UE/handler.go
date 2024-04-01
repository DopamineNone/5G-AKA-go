package main

import (
	"_5gAKA_go"
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"context"
	"fmt"
	"log"
	"time"
)

var (
	seafClient seaf.Client
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context) (resp string, err error) {
	// Init Authentication
	ki, op, snName, _ := InitForUE()
	opc := _5gAKA_go.MilenageGenOpc(ki, op)

	// Send SUPI and SN_Name to SEAF, get response, and update logs
	SUPI := GenerateSUPI()
	SUCI := GenerateSUCI(SUPI)
	authReq, err := seafClient.Authenticate(context.Background(), SUCI+snName)
	log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and SN_name to SEAF")
	if err != nil {
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
		return "", err
	}
	log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive auth-request from SEAF.")

	// Calculate *RES and check
	randNum, AUTN := authReq[:32], authReq[32:]
	sqnAK, amf, xMacA := ResolveAUTN(AUTN)

	res, ck, ik, ak := _5gAKA_go.MilenageF2345(ki, opc, randNum)

	xSqn := _5gAKA_go.LogicalXOR(ak, sqnAK)
	macA, _ := _5gAKA_go.MilenageF1(ki, opc, randNum, xSqn, amf)
	if CheckMac(xMacA, macA) == 1 {
		P0 := snName
		L0 := fmt.Sprintf("%x", len(snName))
		resStar := GenerateResStar(ck, ik, P0, L0, randNum, res)

		// Send auth-response to SEAF and update logs
		_, err := seafClient.Authenticate(context.Background(), resStar)
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send res* to SEAF. Value:" + resStar)
		if err != nil {
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
			return "", err
		} else {
			// Authentication passed!
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Authentication passed successfully!")
			return "Authentication passed successfully!", nil
		}
	}
	return "", fmt.Errorf("Authentication failed: Unable to pass auth-response")
}
