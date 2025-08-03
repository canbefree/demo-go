package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"39.106.55.254:9876"})),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	if err != nil {
		panic(err)
	}
	defer p.Shutdown()

	msg := primitive.NewMessage("TestTopic", []byte("Hello RocketMQ from Go!"))
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Printf("Send error: %s\n", err)
	} else {
		fmt.Printf("Send success: result=%s\n", res.String())
	}
}
