package kafka

import (
	"fmt"

	cluster "github.com/bsm/sarama-cluster"
)

// NewClusterConsumers create a lot of cluster consumer
func NewClusterConsumers(fileConfig string) ([]*ClusterConsumer, error) {
	conf, err := UnmarshalViaFile(fileConfig)
	if err != nil {
		return nil, err
	}

	consumers := make([]*ClusterConsumer, len(conf.Consumers))
	for i, c := range conf.Consumers {
		config := cluster.NewConfig()
		if c.SASL.Enabled {
			config.Net.SASL.Enable = c.SASL.Enabled
			config.Net.SASL.User = c.SASL.User
			config.Net.SASL.Password = c.SASL.Password
		}

		consumer := &ClusterConsumer{
			Brokers: c.Brokers,
			GroupID: c.GroupID,
			Topics:  c.Topics,
			Config:  config,
			OnErrorFunc: func(err error) {
				fmt.Printf("error: %v", err)
			},
		}

		consumers[i] = consumer
	}

	return consumers, nil
}
