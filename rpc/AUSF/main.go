package main

import (
	ausf "_5gAKA_go/kitex_gen/_5gAKA_go/AUSF/protocolservice"
	udm "_5gAKA_go/kitex_gen/_5gAKA_go/UDM/protocolservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
	host    string = "localhost"
	port    string = "8003"
	hostUDM string = "localhost"
	portUDM string = "8004"
	logPath string = "../../log/AUSF.log"
)

func main() {
	// Set global log output
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	// Build connection with UDM
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
