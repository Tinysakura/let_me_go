package proto

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{pub:pubsub.NewPublisher(2 * time.Second, 10)}
}

func (pub *PubsubService)Publish(ctx context.Context, in *String) (*String, error) {
	pub.pub.Publish(in.Value)
	return &String{}, nil
}

func (pub *PubsubService) Subscribe(args *String, stream PubsubService_SubscribeServer) error {
	ch := pub.pub.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			// 只订阅in入参传入的值开头的消息
			if strings.HasPrefix(s, args.Value) {
				return true
			}
		}

		return false
	})

	// 将接收的到的订阅结果发送到订阅客户端
	for v := range ch {
		stream.Send(&String{Value:v.(string)})
	}

	return nil
}