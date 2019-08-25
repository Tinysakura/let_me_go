package main

import (
	"google.golang.org/grpc"
	proto "let_me_go/protobuf"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterPubsubServiceServer(grpcServer, proto.NewPubsubService())

	listen, e := net.Listen("tcp", "localhost:8888")

	if e != nil {
		log.Fatal(e)
	}

	grpcServer.Serve(listen)
}
