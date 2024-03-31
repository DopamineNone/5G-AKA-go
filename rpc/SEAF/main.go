package main

import (
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	host     string = "localhost"
	port     string = "8002"
	ausfHost string = "localhost"
	ausfPort string = "8003"
)

func main() {
	fmt.Println("SEAF:")
	addr, _ := net.ResolveTCPAddr("tcp", host+":"+port)
	svr := seaf.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

}
