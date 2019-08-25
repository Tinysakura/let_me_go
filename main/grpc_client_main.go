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

	client := proto.NewHelloServiceClient(dialConn)
	result, _ := client.Hello(context.Background(), &proto.String{Value: "Hello GRPC"})

	fmt.Println(result.Value)
}
