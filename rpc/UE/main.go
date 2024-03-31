package main

import (
	ue "_5gAKA_go/kitex_gen/_5gAKA_go/UE/protocolservice"
	"log"
)

func main() {
	svr := ue.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
