package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	seaf "_5gAKA_go/kitex_gen/_5gAKA_go/SEAF/protocolservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"io"
	"log"
	"net"
	"os"
)

var (
	hostSEAF string = "localhost"
	portSEAF string = "8002"
	hostAUSF string = "localhost"
	portAUSF string = "8003"
	logPath  string = "../../log/SEAF.log"
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

	// Build connection with AUSF
	var err error
	ausfClient, err = ausf.NewClient("_5gAKA_go.AUSF", client.WithHostPorts(hostAUSF+":"+portAUSF))

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("SEAF:")
	addr, _ := net.ResolveTCPAddr("tcp", hostSEAF+":"+portSEAF)
	svr := seaf.NewServer(new(ProtocolServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
		return
	}

}
