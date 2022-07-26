package main

import "context"

const (
	TopicName  = "firstTopic"
	BrokerAddr = "localhost:9092"
)

func main() {
	ctx := context.Background()

	go Produce(ctx)
	Consume(ctx)
}
