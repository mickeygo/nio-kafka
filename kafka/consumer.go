package kafka

import (
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

	// OnNotificationFunc
	OnNotificationFunc func(notification *cluster.Notification)

	// OnSuccessFunc 当成功接收到消息后会触发该函数
	OnSuccessFunc func(c *cluster.Consumer, m *sarama.ConsumerMessage) bool

	// OnErrorFunc
	OnErrorFunc func(err error)

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

	for {
		select {
		case msg, ok := <-c.consumer.Messages():
			if ok {
				sarama.Logger.Printf("%v\t %s/%d/%d \t %s\t \n", time.Now(), msg.Topic, msg.Partition, msg.Offset, msg.Value)
				if c.OnSuccessFunc != nil {
					c.OnSuccessFunc(c.consumer, msg)
				}

				c.consumer.MarkOffset(msg, "") // ack, ensure the message have been consumered
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
