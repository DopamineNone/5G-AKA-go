package main

// AUSF端 接收SEAF转发的来自UE的SUCI + SN_name，并将其发送给UDM
// 接收UDM发送来的认证向量5G HE_AV，据此存储xRes*并计算hxRes*，计算5G SE_AV并发送给SEAF
// 接收SEAF发送来的由UE生成的Res*，并将其与之前存储的xRes*比较，相同则认证成功

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

var (
	P0       string
	L0       string
	xResStar string
)

// GenerateHxResStar Generate hxRes* using xRes* by hash algorithm sha-256 after receiving HE_AV||SUPI from UDM.
func GenerateHxResStar(randNum, xResStar string) string {
	s := []byte(randNum + xResStar)
	h := sha256.New()
	h.Write(s)
	tmp := hex.EncodeToString(h.Sum(nil))
	hxResStar := tmp[32:]

	return hxResStar
}

// GenerateKseaf Generate K_seaf after receiving HE_AV||SUPI from UDM.
func GenerateKseaf(kAusf, P0, L0 string) string {
	// P0 snName(serving network name)
	// L0 length of P0
	s := []byte("6C" + P0 + L0)
	h := hmac.New(sha256.New, []byte(kAusf))
	h.Write(s)
	kSeaf := hex.EncodeToString(h.Sum(nil))

	return kSeaf
}

// ResolveDataFromUDM Get HE_AV and SUPI by resolving the data(HE_AV||SUPI) received from UDM.
func ResolveDataFromUDM(data string) (string, string) {
	heAV := data[:160]
	SUPI := data[160:]
	return heAV, SUPI
}

// ResolveDataFromSEAF Get SUCI and snName by resolving the data(SUCI||snName) received from SEAF.
func ResolveDataFromSEAF(data string) (string, string) {
	SUCI := data[:21]
	snName := data[21:]
	return SUCI, snName
}

// ResolveHEAV Get (randNum,AUTN,hxRes*,K_seaf) by resolving HE_AV received from UDM.
func ResolveHEAV(heAV string) (string, string, string, string) {
	randNum := heAV[:32]
	AUTN := heAV[32:64]
	xResStarr := heAV[64:96]
	kAusf := heAV[96:]
	return randNum, AUTN, xResStarr, kAusf
}

// GenerateSEAV Compute SE_AV, which will be sent to SEAF
func GenerateSEAV(randNum, AUTN, hxResStar, kSeaf string) string {
	return randNum + AUTN + hxResStar + kSeaf
}
