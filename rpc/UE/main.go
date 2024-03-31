package main

import (
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	ue "_5gAKA_go/kitex_gen/_5gAKA_go/UE/protocolservice"
	"fmt"
	"log"
)

func main() {
	fmt.Println("UE")
	svr := ue.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

	seafClient, err = seaf.NewClient("_5gAKA_go.SEAF")

	if err != nil {
		log.Println(err.Error())
	}

}
