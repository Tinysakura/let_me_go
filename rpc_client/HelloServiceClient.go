package rpc_client

import (
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
	client.client.Call(rpc_server.Hello_Service_Name + ".Hello", request, response)

	return nil
}
