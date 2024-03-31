package main

import (
	"_5gAKA_go"
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"context"
	"fmt"
	"os"
	"time"
)

var (
	logPath    string = "../../log/UE.log"
	seafClient seaf.Client
	failMsg    string = "Authentication failed: Server error."
	successMsg string = "Authentication was completed successfully!"
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context) (resp string, err error) {
	// Init Authentication
	ki, op, snName, _ := InitForUE()
	opc := _5gAKA_go.MilenageGenOpc(ki, op)

	// Load log file
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Send SUPI and SN_Name to SEAF, get response, and update logs
	SUPI := GenerateSUPI()
	SUCI := GenerateSUCI(SUPI)
	authReq, err := seafClient.Authenticate(context.Background(), SUCI+snName)
	_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and SN_name to SEAF.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and SN_name to SEAF")
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Failed to receive reponse from SEAF")
		return failMsg, err
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive auth-request from SEAF.")
	_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive auth-request from SEAF.")

	randNum, AUTN := authReq[:32], authReq[32:]
	sqnAK, amf, xMacA := ResolveAUTN(AUTN)

	res, ck, ik, ak := _5gAKA_go.MilenageF2345(ki, opc, randNum)

	xSqn := _5gAKA_go.LogicalXOR(ak, sqnAK)
	macA, _ := _5gAKA_go.MilenageF1(ki, opc, randNum, xSqn, amf)
	if CheckMac(xMacA, macA) == 1 {
		P0 := snName
		L0 := fmt.Sprintf("%x", len(snName))
		resStar := GenerateResStar(ck, ik, P0, L0, randNum, res)

		_, err = seafClient.Authenticate(context.Background(), resStar)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send res* to SEAF. Value:" + resStar)
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send res* to SEAF. Value:" + resStar)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Failed to receive reponse from SEAF")
			_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Failed to receive reponse from SEAF")
			return failMsg, err
		}
	}
	return successMsg, nil
}
