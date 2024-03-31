package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	hostSEAF string = "localhost"
	portSEAF string = "8002"
	hostAUSF string = "localhost"
	portAUSF string = "8003"
)

func main() {
	// First: build connection with AUSF
	var err error
	ausfClient, err = ausf.NewClient("_5gAKA_go.AUSF", client.WithHostPorts(hostAUSF+":"+portAUSF))

	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("SEAF:")
	addr, _ := net.ResolveTCPAddr("tcp", hostSEAF+":"+portSEAF)
	svr := seaf.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

}
