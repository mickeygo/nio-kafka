package goka

import (
	"time"

	"github.com/Shopify/sarama"
)

// SyncProducer producer
type SyncProducer struct {
	Brokers []string
	Topic   string
	Config  *sarama.Config

	OnPostedSuccessFunc func(msg *sarama.ProducerMessage)
	OnErrorFunc         func(err error)

	producer sarama.SyncProducer
}

// SendMessage send message
func (p *SyncProducer) SendMessage(value string) error {
	if err := p.newSyncProducer(); err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.ByteEncoder(value),
	}

	p.syncProducer(msg)

	return nil
}

// SendProducerMessage send producer message
func (p *SyncProducer) SendProducerMessage(msg *sarama.ProducerMessage) error {
	if err := p.newSyncProducer(); err != nil {
		return err
	}

	defer p.producer.Close()
	p.syncProducer(msg)

	return nil
}

func (p *SyncProducer) newSyncProducer() error {
	if p.Config == nil {
		p.Config = sarama.NewConfig()
		p.Config.Producer.Return.Successes = true
		p.Config.Producer.Timeout = 5 * time.Second
	}

	pd, err := sarama.NewSyncProducer(p.Brokers, p.Config)
	if err != nil {
		return err
	}

	p.producer = pd

	return nil
}

func (p *SyncProducer) syncProducer(msg *sarama.ProducerMessage) {
	if p.producer == nil {
		return
	}

	if _, _, err := p.producer.SendMessage(msg); err != nil {
		p.OnErrorFunc(err)
	}

	if p.OnPostedSuccessFunc != nil {
		p.OnPostedSuccessFunc(msg)
	}
}
