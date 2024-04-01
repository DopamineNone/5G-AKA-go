package main

import (
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"github.com/cloudwego/kitex/server"
	"io"
	"log"
	"net"
	"os"
)

var (
	hostUDM string = "localhost"
	portUDM string = "8004"
	logPath string = "../../log/UDM.log"
)

func main() {
	// Load log file
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Set log output
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	log.Println("UDM:")
	addr, _ := net.ResolveTCPAddr("tcp", hostUDM+":"+portUDM)
	svr := udm.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
