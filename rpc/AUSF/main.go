package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"fmt"
	"log"
)

func main() {
	fmt.Println("AUSF:")
	svr := ausf.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

	udmClient, err = udm.NewClient("_5gAKA_go.UDM")

	if err != nil {
		log.Println(err.Error())
	}
}
