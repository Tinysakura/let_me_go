package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "let_me_go/protobuf"
)

func main() {
	dialConn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())

	defer dialConn.Close()

	client := proto.NewPubsubServiceClient(dialConn)

	stream, _ := client.Subscribe(context.Background(), &proto.String{Value: "ios"})

	for  {
		recv, _ := stream.Recv()
		fmt.Println("接收到订阅消息:" + recv.Value)
	}
}
