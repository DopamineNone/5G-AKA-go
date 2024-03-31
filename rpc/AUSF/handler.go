package main

import (
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"context"
	"fmt"
	"os"
	"time"
)

var (
	logPath   string = "../../log/AUSF.log"
	udmClient udm.Client
)

// ProtocolServiceImpl implements the last service interface defined in the IDL.
type ProtocolServiceImpl struct{}

// Authenticate implements the ProtocolServiceImpl interface.
func (s *ProtocolServiceImpl) Authenticate(ctx context.Context, data string) (resp string, err error) {

	// TODO: Your code here...
	// 根据消息长度做出不同反应
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	length := len(data)

	if length == 30 {
		// 消息长度为30，则为SEAF发来的SUCI + SN_name {SUCI---21, SN_name---9}
		// 把SUCI + SN_name发送给UDM
		_, snName := ResolveDataFromSEAF(data)
		P0 = snName
		L0 = fmt.Sprintf("%x", len(snName))

		resp, err := udmClient.Authenticate(context.Background(), data)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI nd SN_name to UDM.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send SUCI nd SN_name to UDM.")
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "UDM Authentication failed.")
			_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "UDM Authentication failed.")
			return "UDM Authentication failed.", fmt.Errorf("UDM Authentication failed.")
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive 5G HE_AV and SUPI from UDM.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Receive 5G HE_AV and SUPI from UDM.")

		heAV, SUPI := ResolveDataFromUDM(resp)
		randNum, AUTN, _, kAusf := ResolveHEAV(heAV)

		_, _, xResStar, _ = ResolveHEAV(heAV)
		kSeaf := GenerateKseaf(kAusf, P0, L0)
		hxResStar := GenerateHxResStar(randNum, xResStar)

		seAV := GenerateSEAV(randNum, AUTN, hxResStar, kSeaf)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send 5G SE_AV and SUPI to SEAF.")
		_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send 5G SE_AV and SUPI to SEAF.")

		return seAV + SUPI, nil
	} else if length == 32 {
		// 消息长度为32，则为SEAF发送来的Res*  {Res*---32}
		// 将SEAF发送来的Res*和之前从UDM接收的hRes*进行比较，相同则说明认证成功。
		if data == xResStar {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send Authentication Response to SEAF.")
			_, _ = file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "  " + "Send Authentication Response to SEAF.")
			return "AUSF Authentication passed", nil
		}
	}
	return
}
