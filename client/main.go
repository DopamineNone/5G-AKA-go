package main

import (
	ue "_5gAKA_go/kitex_gen/_5gAKA_go/UE/protocolservice"
	"context"
	"fmt"
)

func main() {
	cli, err := ue.NewClient("_5gAKA_go.UE")
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
		return
	}
	err = cli.Authenticate(context.Background())
	if err == nil {
		fmt.Println("Authentication complete successfully!")
	} else {
		fmt.Println("Error occurred: ", err.Error())
	}
}
