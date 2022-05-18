package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

// SyncKafkaProducer 同步kafka生产者
type SyncKafkaProducer struct {
	KafkaConfig *KafkaConfig
}

func (k *SyncKafkaProducer) NewKafkaProducer(address []string, topic string, duration time.Duration) {
	if len(address) == 0 {
		panic(addressError)
	}
	k.KafkaConfig = NewProducerByMessage(address, topic, duration)
}

func (k *SyncKafkaProducer) Send(value []byte) {
	if k == nil || k.KafkaConfig == nil || value == nil {
		return
	}
	syncProducer, err := sarama.NewSyncProducer(k.KafkaConfig.addressList, k.KafkaConfig.config)
	if err != nil {
		panic(newSyncProducerError)
	}

	defer syncProducer.Close()
	msg := &sarama.ProducerMessage{
		Topic: k.KafkaConfig.topic,
		Value: sarama.ByteEncoder(value),
	}
	_, _, err = syncProducer.SendMessage(msg)
	if err != nil {
		panic(sendMessageError)
	}
}
