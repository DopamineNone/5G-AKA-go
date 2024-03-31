package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	"log"
)

func main() {
	svr := ausf.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
