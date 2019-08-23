package main

import (
	"fmt"
	"let_me_go/rpc_client"
)

func main() {
	ip := "127.0.0.1:8888"
	helloServiceClient := rpc_client.Dial(ip)

	response := ""
	e := helloServiceClient.Hello("", &response)

	if e == nil {
		fmt.Println(response)
	}
}
