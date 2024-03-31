package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"time"
)

var (
	host    string = "localhost"
	port    string = "8003"
	hostUDM string = "localhost"
	portUDM string = "8004"
)

func main() {
	// First: build connection with UDM
	var err error
	udmClient, err = udm.NewClient("_5gAKA_go.UDM", client.WithHostPorts(hostUDM+":"+portUDM))

	if err != nil {
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
		return
	}

	log.Println("AUSF:")
	addr, _ := net.ResolveTCPAddr("tcp", host+":"+port)
	svr := ausf.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + err.Error())
		return
	}
}
