package proto

import (
	"context"
	"fmt"
)

type HelloServiceImpl struct {

}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello grpc"}

	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for  {
		args, _ := stream.Recv()

		if args.Value == "end" {
			fmt.Println("stop connection")
			return nil
		} else {
			fmt.Println(args.Value)
		}

		e := stream.Send(&String{Value: "你说我听着呢"})

		if e != nil {
			return e
		}
	}
}
