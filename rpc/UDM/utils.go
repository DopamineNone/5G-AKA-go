package main

// UDM端 接收AUSF发来的SUCI和SN_name，解密得到SUPI
// 生成5G HE_AV并发送给AUSF端

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

// GetSUPI Get SUPI by decrypting SUCI.
func GetSUPI(SUCI string) string {
	mcc := SUCI[1:4]
	mnc := SUCI[4:6]
	msin := SUCI[11:]
	return mcc + mnc + msin
}

// GenerateRand Generate a 128-bit rand number.
func GenerateRand() (result string) {
	chars := "0123456789abcdef"
	for i := 1; i < 33; i++ {
		result += string(chars[rand.Intn(16)])
	}
	return result
}

// GenerateKausf Generate K_ausf.
func GenerateKausf(key, P0, L0, P1, L1 string) string {
	appSecret := []byte(key)
	s := []byte("6A" + P0 + L0 + P1 + L1)
	h := hmac.New(sha256.New, appSecret)
	h.Write(s)
	tmp := hex.EncodeToString(h.Sum(nil))
	ckNew := tmp[:32]
	ikNew := tmp[32:]
	keyNew := ckNew + ikNew
	h1 := hmac.New(sha256.New, []byte(keyNew))
	h1.Write(s)
	kAusf := hex.EncodeToString(h1.Sum(nil))
	return kAusf
}

// GenerateXResStar Generate xRes*.
func GenerateXResStar(key, P0, L0, P1, L1, P2, L2 string) string {
	appSecret := []byte(key)
	s := []byte("6B" + P0 + L0 + P1 + L1 + P2 + L2)
	h := hmac.New(sha256.New, appSecret)
	h.Write(s)
	tmp := hex.EncodeToString(h.Sum(nil))
	xResStar := tmp[32:]
	return xResStar
}

func InitForUDM() (string, string, string, string, string) {
	ki := "000000012449900000000010123456d8"
	randNum := GenerateRand()
	sqn := "1234567888d8"
	amf := "8d00"
	op := "cda0c2852846d8eb63a387051cdd1fa5"
	return ki, randNum, sqn, amf, op
}
