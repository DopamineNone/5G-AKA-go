package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	host    string = "localhost"
	port    string = "8003"
	udmHost string = "localhost"
	udmPort string = "8004"
)

func main() {
	fmt.Println("AUSF:")
	addr, _ := net.ResolveTCPAddr("tcp", host+":"+port)
	svr := ausf.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

}
