package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"fmt"
	"log"
)

func main() {
	fmt.Println("SEAF:")
	svr := seaf.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

	ausfClient, err = ausf.NewClient("_5gAKA_go.AUSF")

	if err != nil {
		log.Println(err.Error())
	}
}
