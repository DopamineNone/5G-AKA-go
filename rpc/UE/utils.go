package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	//"os"
	"strconv"
)

// UE端 产生SUPI并处理得到SUCI，把SUCI以及SN_name发送给SN
// 接收SEAF发送来的5G SE_AV，提取其中的各种值(randNum,AUTN),AUTN=xSQN^AK||amf||x_macA,处理后利用milenage算法计算得到macA
// 比较两个mac值是否相同，若相同则生成Res，进一步生成Res*并将其发送给SEAF

// GenerateSUPI Generate Subscription Permanent Identifier(SUPI).
func GenerateSUPI() string {
	//imsi := "46000"，IMSI是SUPI的一种类型，IMSI=MCC||MNC||MSIN, 3+2+10位十进制数字
	SUPI := "46000"
	for i := 0; i < 10; i++ {
		SUPI = SUPI + strconv.Itoa(rand.Intn(10))
	}
	return SUPI
}

// GenerateSUCI Generate Subscription Concealed Identifier(SUCI). (use null scheme)
func GenerateSUCI(SUPI string) string {
	// 对SUPI使用空保护策略
	// SUCI=SUPI类型取值(0表示imsi)||归属网络标识符(mcc+mnc)||路由标识符||SUPI保护算法ID(0表示null scheme)||归属网络公钥||msin
	mcc := SUPI[:3]
	mnc := SUPI[3:5]
	msin := SUPI[5:]
	SUCI := "0" + mcc + mnc + "678" + "0" + "0" + msin
	return SUCI
}

func ResolveAUTN(AUTN string) (string, string, string) {
	sqnAK := AUTN[:12]
	amf := AUTN[12:16]
	mac := AUTN[16:]
	return sqnAK, amf, mac
}

func CheckMac(xMacA, MacA string) int {
	if xMacA == MacA {
		return 1
	} else {
		return 0
	}
}

func GenerateResStar(ck, ik, P0, L0, rand, res string) string {
	key := []byte(ck + ik)
	P1 := rand
	L1 := fmt.Sprintf("%x", len(P1))
	P2 := res
	L2 := fmt.Sprintf("%x", len(P2))
	s := []byte("6B" + P0 + L0 + P1 + L1 + P2 + L2)
	h := hmac.New(sha256.New, key)
	h.Write(s)
	resStar := hex.EncodeToString(h.Sum(nil))[32:]
	return resStar
}

func InitForUE() (string, string, string, string) {
	ki := "000000012449900000000010123456d8"
	op := "cda0c2852846d8eb63a387051cdd1fa5"
	//global sn_name
	snName := "123456789"
	sqnMax := "100000000000000000000000"
	return ki, op, snName, sqnMax
}
