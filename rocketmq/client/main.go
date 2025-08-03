package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test_consumer_group"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"39.106.55.254:9876"})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Subscribe("TestTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("Received message: %s\n", msg.Body)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Consumer started. Waiting for messages...")
	select {} // 阻塞主线程
}
