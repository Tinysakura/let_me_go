package rpc_server

import (
	"net"
	"net/rpc"
	"time"
)

const Hello_Service_Name = "HelloService"

type HelloService interface {
	Hello (request string, response *string) error
}

type Hello struct {

}

func (hello *Hello) Hello(request string, response *string) error {
	// 人工阻塞
	time.Sleep(time.Duration(3) * time.Second)
	*response = "Hello go rpc"

	return nil
}

func Register(ip string) {
	listener, _ := net.Listen("tcp", ip)
	for {
		conn, _ := listener.Accept()

		rpc.RegisterName(Hello_Service_Name, new(Hello))
		rpc.ServeConn(conn)
	}
}
