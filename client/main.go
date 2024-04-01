package main

import (
	ue "_5gAKA_go/kitex_gen/_5gAKA_go/UE/protocolservice"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
)

var (
	host string = "localhost"
	port string = "8001"
)

func main() {
	cli, err := ue.NewClient("_5gAKA_go.UE", client.WithHostPorts(host+":"+port))
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
		return
	}
	resp, err := cli.Authenticate(context.Background())
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println("Error occurred: ", err.Error())
	}
}
