package main

import (
	"fmt"
	"github.com/Maksim-Gr/kafkaConsumer/pkg/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	fmt.Println(config.Environment)

}
