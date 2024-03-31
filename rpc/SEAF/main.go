package main

import (
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"log"
)

func main() {
	svr := seaf.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
