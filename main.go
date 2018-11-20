package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mickeygo/nio-kafka/goka"
)

func main() {
	consumers, err := goka.NewClusterConsumers("config.yml")
	if err != nil {
		log.Fatalf("read the config file error. %v", err)
		os.Exit(1)
	}

	for _, consumer := range consumers {
		go func(c *goka.ClusterConsumer) {
			defer c.Close()
			fmt.Printf("start listen the kafka server %s, topics: %v ... \n", c.Brokers, c.Topics)
			c.Poll()
		}(consumer)
	}

	select {}
}
