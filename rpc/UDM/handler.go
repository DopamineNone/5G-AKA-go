package main

import (
	"_5gAKA_go"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	logPath string = "../../log/UDM.log"
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

	// Set log output
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive SUCI and SN_name from AUSF.")

	SUCI, snName := data[:21], data[21:]
	SUPI := GetSUPI(SUCI)

	ki, randNum, sqn, amf, op := InitForUDM()
	opc := _5gAKA_go.MilenageGenOpc(ki, op)

	xRes, ck, ik, AUTN, ak := _5gAKA_go.Milenage(ki, opc, randNum, sqn, amf)

	key := ck + ik
	P0 := snName
	L0 := fmt.Sprintf("%x", len(P0))
	P1 := _5gAKA_go.LogicalXOR(sqn, ak)
	L1 := fmt.Sprintf("%x", len(P1))
	kAusf := GenerateKausf(key, P0, L0, P1, L1)

	P1 = randNum
	L1 = fmt.Sprintf("%x", len(P1))
	P2 := xRes
	L2 := fmt.Sprintf("%x", len(P2))

	xResStar := GenerateXResStar(key, P0, L0, P1, L1, P2, L2)

	heAV := randNum + AUTN + xResStar + kAusf

	log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send 5G HE AV and SUPI to AUSF.")

	return heAV + SUPI, nil
}
