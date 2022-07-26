package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context) {
	i := 0

	//initializing the writer with topic and broker addr
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{BrokerAddr},
		Topic:   TopicName,
		//BatchSize:    2,
		//BatchTimeout: time.Millisecond * 5,
		//RequiredAcks: -1,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("This is message number: " + strconv.Itoa(i)),
		})

		if err != nil {
			panic("Could not write the message " + err.Error())
		}

		//confirmation for message is written
		fmt.Println("Wrote: ", i)
		i++
		if i == 8 {
			break
		}

		//Sleeping so consumer can consume data
		//time.Sleep(time.Millisecond * 100)
	}
}
