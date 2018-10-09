package goka

import (
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// ClusterConsumer a cluster consumer
type ClusterConsumer struct {
	Brokers []string
	Group   string
	Topics  []string
	Config  *cluster.Config

	OnNotificationFunc func(notification *cluster.Notification)
	OnSuccessFunc      func(c *cluster.Consumer, m *sarama.ConsumerMessage)
	OnErrorFunc        func(err error)

	consumer *cluster.Consumer
}

// NewConsumer create a new consumer
func (c *ClusterConsumer) newConsumer() error {
	if c.Config == nil {
		c.Config = cluster.NewConfig()
	}
	c.Config.Consumer.Return.Errors = true
	c.Config.Group.Return.Notifications = true

	cu, err := cluster.NewConsumer(c.Brokers, c.Group, c.Topics, c.Config)
	if err != nil {
		return err
	}

	c.consumer = cu

	return nil
}

// Poll listen the message
func (c *ClusterConsumer) Poll() error {
	if err := c.newConsumer(); err != nil {
		return err
	}

	// listen
	for {
		select {
		case msg, more := <-c.consumer.Messages():
			if more {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t/%v\n", msg.Topic, msg.Partition, msg.Offset, msg.Value, time.Now())
				if c.OnSuccessFunc != nil {
					c.OnSuccessFunc(c.consumer, msg)
				}
			}
		case ntf, more := <-c.consumer.Notifications():
			if more {
				sarama.Logger.Printf("Rebalanced: %+v\n", ntf)
				if c.OnNotificationFunc != nil {
					c.OnNotificationFunc(ntf)
				}
			}
		case err, more := <-c.consumer.Errors():
			if more {
				sarama.Logger.Printf("Error: %s\n", err.Error())
				if c.OnErrorFunc != nil {
					c.OnErrorFunc(err)
				}
			}
		}
	}
}

// Close close the consumer
func (c *ClusterConsumer) Close() {
	if c.consumer != nil {
		c.consumer.Close()
	}
}
