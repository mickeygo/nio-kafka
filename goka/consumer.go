package goka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// ClusterConsumer a cluster consumer
type ClusterConsumer struct {
	Brokers []string
	GroupID string
	Topics  []string
	Config  *cluster.Config

	OnNotificationFunc func(notification *cluster.Notification)
	OnSuccessFunc      func(c *cluster.Consumer, m *sarama.ConsumerMessage) bool
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
	c.Config.Consumer.Offsets.Initial = sarama.OffsetNewest //

	cu, err := cluster.NewConsumer(c.Brokers, c.GroupID, c.Topics, c.Config)
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
		case msg, ok := <-c.consumer.Messages():
			if ok {
				fmt.Printf("%v\t %s/%d/%d \t %s\t \n", time.Now(), msg.Topic, msg.Partition, msg.Offset, msg.Value)

				if c.OnSuccessFunc != nil {
					c.OnSuccessFunc(c.consumer, msg)
				}

				c.consumer.MarkOffset(msg, "") // 确认消息已成功被消费
			}
		case ntf, ok := <-c.consumer.Notifications():
			if ok {
				sarama.Logger.Printf("Rebalanced: %+v\n", ntf)
				if c.OnNotificationFunc != nil {
					c.OnNotificationFunc(ntf)
				}
			}
		case err, ok := <-c.consumer.Errors():
			if ok {
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
