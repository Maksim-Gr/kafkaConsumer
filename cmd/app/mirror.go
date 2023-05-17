package main

import (
	"fmt"
	"github.com/Maksim-Gr/kafkaConsumer/pkg/util"
	"log"
)

var (
	brokers = ""
	group   = ""
	topic   = ""
)

func main() {
	keepRunning := true
	for keepRunning {
		select {}
	}

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(config.Environment)

}
