package main

import (
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"context"
	"fmt"
	"log"
	"time"
)

var (
	udmClient udm.Client
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context, data string) (resp string, err error) {
	// 根据消息长度做出不同反应

	log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive SUCI and SN_name from SEAF.")
	length := len(data)

	if length == 30 {
		// 消息长度为30，则为SEAF发来的SUCI + SN_name {SUCI---21, SN_name---9}
		// 把SUCI + SN_name发送给UDM
		_, snName := ResolveDataFromSEAF(data)
		P0 = snName
		L0 = fmt.Sprintf("%x", len(snName))

		resp, err := udmClient.Authenticate(context.Background(), data)
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI and SN_name to UDM.")
		if err != nil {
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
			return "", err
		}
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive 5G HE_AV and SUPI from UDM.")

		heAV, SUPI := ResolveDataFromUDM(resp)
		randNum, AUTN, _, kAusf := ResolveHEAV(heAV)

		_, _, xResStar, _ = ResolveHEAV(heAV)
		kSeaf := GenerateKseaf(kAusf, P0, L0)
		hxResStar := GenerateHxResStar(randNum, xResStar)

		seAV := GenerateSEAV(randNum, AUTN, hxResStar, kSeaf)
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send 5G SE_AV and SUPI to SEAF.")

		return seAV + SUPI, nil
	} else if length == 32 {
		// 消息长度为32，则为SEAF发送来的Res*  {Res*---32}
		// 将SEAF发送来的Res*和之前从UDM接收的hRes*进行比较，相同则说明认证成功。
		if data == xResStar {
			log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "AUSF Authentication Passed! Send Authentication Response to SEAF.")
			return "AUSF Authentication passed", nil
		}
	} else {
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "AUSF Authentication failed: Unable to process unknown data")
	}
	return
}
