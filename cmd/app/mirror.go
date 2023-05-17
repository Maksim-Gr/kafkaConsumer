package main

import (
	"fmt"
	"github.com/Maksim-Gr/kafkaConsumer/pkg/util"
	"gopkg.in/Shopify/sarama.v1"
	"log"
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

func main() {
	keepRunning := true
	log.Println("starting mirror consumer")

	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	for keepRunning {
		select {}
	}

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(config.Environment)

}
