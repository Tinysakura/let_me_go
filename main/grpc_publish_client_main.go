package main

import (
	"context"
	"google.golang.org/grpc"
	proto "let_me_go/protobuf"
	"time"
)

func main() {
	dialConn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())

	defer dialConn.Close()

	client := proto.NewPubsubServiceClient(dialConn)

	time.Sleep(time.Duration(5) * time.Second)

	client.Publish(context.Background(), &proto.String{Value: "安卓也不错"})

	client.Publish(context.Background(), &proto.String{Value: "ios是最棒滴"})

}
