package rpc_client

import (
	"fmt"
	"let_me_go/rpc_server"
	"net/rpc"
)

type HelloServiceClient struct {
	client *rpc.Client
}

func Dial(ip string) *HelloServiceClient {
	rpcClient, _ := rpc.Dial("tcp", ip)
	return &HelloServiceClient{client:rpcClient}
}

func (client *HelloServiceClient)Hello(request string, response *string) error {

	// client.client.Call(rpc_server.Hello_Service_Name + ".Hello", request, response)
	// 异步化调用改造
	done := make(chan *rpc.Call, 1)
	call := client.client.Go(rpc_server.Hello_Service_Name+".Hello", request, response, done)

	fmt.Println("hh我没被阻塞")

	<- call.Done
	reply := call.Reply.(*string)
	fmt.Println(*reply)
	return nil
}
