package main

import (
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	ue "_5gAKA_go/kitex_gen/_5gAKA_go/UE/protocolservice"
	"github.com/cloudwego/kitex/client"
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
	// First: build connection with SEAF
	var err error
	seafClient, err = seaf.NewClient("_5gAKA_go.SEAF", client.WithHostPorts(seafHost+":"+seafPort))

	if err != nil {
		log.Println(err.Error())
		return
	}

	// Second: run UE server
	log.Println("UE:")

	addr, _ := net.ResolveTCPAddr("tcp", host+":"+port)
	svr := ue.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}
}
