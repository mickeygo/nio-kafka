package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mickeygo/nio-kafka/kafka"
)

func main() {
	consumers, err := kafka.NewClusterConsumers("config.yml")
	if err != nil {
		log.Fatalf("read the config file error. %v", err)
		os.Exit(1)
	}

	for _, consumer := range consumers {
		go func(c *kafka.ClusterConsumer) {
			defer c.Close()
			fmt.Printf("start listen the kafka server %s, topics: %v ... \n", c.Brokers, c.Topics)
			c.Poll()
		}(consumer)
	}

	select {}
}
