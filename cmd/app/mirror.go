package main

import (
	"context"
	"fmt"
	"github.com/Maksim-Gr/kafkaConsumer/pkg/util"
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"time"
)

var (
	brokers = ""
	group   = ""
	topic   = ""
)

// Consumer represents consumer
type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			log.Printf("message value: %s", string(message.Value))
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}

func main() {
	keepRunning := true
	log.Println("starting mirror consumer")

	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerConfig.Consumer.MaxProcessingTime = 10 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, consumerConfig)
	if err != nil {
		log.Panicf("error creating consumer group clinet: %v", err)
	}

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(config.Environment)

}
