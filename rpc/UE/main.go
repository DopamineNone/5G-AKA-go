package main

import (
	ue "_5gAKA_go/kitex_gen/_5gAKA_go/UE/protocolservice"
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	host     string = "localhost"
	port     string = "8001"
	seafHost string = "localhost"
	seafPort string = "8002"
)

func main() {
	fmt.Println("UE:")

	addr, _ := net.ResolveTCPAddr("tcp", host+":"+port)
	svr := ue.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}
}
