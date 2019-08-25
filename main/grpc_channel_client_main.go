package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	proto "let_me_go/protobuf"
	"time"
)

func main() {
	dialConn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())

	defer dialConn.Close()

	client := proto.NewHelloServiceClient(dialConn)

	stream, _ := client.Channel(context.Background())

	go func() {
		for  {
			stream.Send(&proto.String{Value:"诶呀我说呀"})
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()

	for  {
		recv, e := stream.Recv()
		if e != nil {
			if e == io.EOF {
				break
			}
		}

		fmt.Println(recv.Value)
	}
}