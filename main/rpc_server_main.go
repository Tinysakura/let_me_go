package main

import "let_me_go/rpc_server"

func main() {
	ip := "127.0.0.1:8888"
	rpc_server.Register(ip)
}
