package main

import (
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	host string = "localhost"
	port string = "8004"
)

func main() {
	fmt.Println("UDM:")
	addr, _ := net.ResolveTCPAddr("tcp", host+":"+port)
	svr := udm.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
