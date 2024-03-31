package main

import (
	base "_5gAKA_go/kitex_gen/_5gAKA_go/base/protocolservice"
	"log"
)

func main() {
	svr := base.NewServer(new(ProtocolServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
