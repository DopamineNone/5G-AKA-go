package main

import (
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	hostUDM string = "localhost"
	portUDM string = "8004"
)

func main() {
	fmt.Println("UDM:")
	addr, _ := net.ResolveTCPAddr("tcp", hostUDM+":"+portUDM)
	svr := udm.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
