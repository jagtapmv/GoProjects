package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context) {
	//Initialize reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Topic:   TopicName,
		Brokers: []string{BrokerAddr},
		GroupID: "first-consumer-group",
		//MinBytes:    1,
		//MaxBytes:    2,
		//MaxWait:     time.Millisecond * 10,
		//StartOffset: kafka.FirstOffset,
	})

	for {
		msg, err := r.ReadMessage(ctx)

		if err != nil {
			panic("Could not read message: " + err.Error())
		}

		//log message after receiving it
		fmt.Println("Received Message: ", string(msg.Value))
	}
}
