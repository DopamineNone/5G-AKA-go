package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// SEAF端 监听SEAF对应端口，根据收到的不同信息做出不同处理
// 接收UE发送的SUCI+SN_name,并转发给AUSF
// 从AUSF端接收SE_AV(randNum||AUTN||hxResStar||kSeaf)
// 向UE端发送randNum||AUTN
// 从UE端接收Res*，并产生哈希值hRes*
// 验证计算的hRes*和从AUSF端接收处理得到的hxRes*是否相同，若相同，把Res*发送给AUSF

var (
	randNum   string
	hxResStar string
)

// GenerateHResStar Generate hRes* after receiving RES* from UE.
func GenerateHResStar(randNum, resStar string) string {
	s := []byte(randNum + resStar)
	h := sha256.New()
	h.Write(s)
	tmp := hex.EncodeToString(h.Sum(nil))
	hResStar := tmp[32:]
	return hResStar
}

// ResolveDataFromAusf Resolve data(SE_AV||SUPI) from AUSF.
func ResolveDataFromAusf(data string) (string, string) {
	AV := data[:160]
	SUPI := data[160:]
	return AV, SUPI
}

// ResolveAV Resolve SE_AV received from AUSF.
func ResolveAV(AV string) (string, string, string, string) {
	randNumber := AV[:32]
	AUTN := AV[32:64]
	hxResStar := AV[64:96]
	kSeaf := AV[96:]
	return randNumber, AUTN, hxResStar, kSeaf
}
