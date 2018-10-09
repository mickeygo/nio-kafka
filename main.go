package main

import (
	"fmt"
	"log"
	"os"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/mickeygo/nio-kafka/goka"
)

func main() {
	conf, err := goka.UnmarshalViaFile("config.yml")
	if err != nil {
		log.Fatalf("read the config file error. %v", err)
		os.Exit(1)
	}

	fmt.Printf("Goka Version: %s \n", conf.Version)

	go func(cfg *goka.Config) {
		for _, c := range cfg.Consumers {
			go func(ccfg *goka.Consumer) {
				config := cluster.NewConfig()
				if ccfg != nil {
					if ccfg.SASL.Enabled {
						config.Net.SASL.Enable = ccfg.SASL.Enabled
						config.Net.SASL.User = ccfg.SASL.User
						config.Net.SASL.Password = ccfg.SASL.Password
					}

					fmt.Printf("connected the kafka [%s] ... \n", ccfg.Name)
					consumer := goka.ClusterConsumer{
						Brokers: ccfg.Brokers,
						Group:   ccfg.Group,
						Topics:  ccfg.Topics,
						Config:  config,
					}
					defer consumer.Close()

					fmt.Printf("start listen the kafka [%s] ... \n", ccfg.Name)
					consumer.Poll()
				}
			}(c)
		}
	}(conf)

	// config := cluster.NewConfig()
	// config.Net.SASL.Enable = true
	// config.Net.SASL.User = "RJLCwcmp"
	// config.Net.SASL.Password = "0DMz8V6Me1gUA9lV"

	// fmt.Fprintln(os.Stdout, "connected the kafka...")

	// consumer := goka.NewConsumer([]string{"10.110.2.66:9092", "10.110.2.69:9092", "10.110.2.72:9092"}, "toolkit", []string{"do-people2-toolkit-dev"}, config)
	// defer consumer.Close()

	// fmt.Fprintln(os.Stdout, "start listen the kafka...")

	// goka.Poll(consumer)
}
