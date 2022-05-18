package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

const (
	// 超时
	kafkaTimeOut = time.Second * 5
	// Sync Async kafka生产者发送信息方式
	Sync  = "Sync"
	Async = "Async"
)

var (
	addressError           = errors.New("kafka address is error")          // 地址错误
	newSyncProducerError   = errors.New("sarama.NewSyncProducer err")      // 创建同步生产者错误
	newAsyncProducerError  = errors.New("sarama.NewAsyncProducer err")     // 创建异步生产者错误
	asyncProducerNullError = errors.New("sarama.NewSyncProducer is null")  // 异步生产者为空
	sendMessageError       = errors.New("send kafka message err")          // 发送信息错误
	modeError              = errors.New("kafka mode error  sync or async") // 模式错误

)

// KafkaConfig kafka生产者
type KafkaConfig struct {
	addressList []string       // 地址列表
	topic       string         // kafka topic
	config      *sarama.Config // kafka配置信息
}

// NewProducer 新建kafka生产者并选择同步方式
// 一般使用异步方式
func NewProducer(address []string, topic string, duration time.Duration, syncOrAsync string) AbsKafkaProducer {
	var producer AbsKafkaProducer
	switch syncOrAsync {
	case Sync:
		producer = &SyncKafkaProducer{}
	case Async:
		producer = &AsyncKafkaProducer{}
	default:
		panic(modeError)
	}
	producer.NewKafkaProducer(address, topic, duration)
	return producer
}

// NewProducerByMessage 创建kafka基本生产者
func NewProducerByMessage(address []string, topic string, duration time.Duration) *KafkaConfig {

	// 根据字符串解析地址列表
	if len(address) < 1 || address[0] == "" {
		panic(addressError)
	}
	// 配置producer参数
	sendConfig := sarama.NewConfig()
	sendConfig.Producer.Return.Successes = true
	sendConfig.Producer.Timeout = kafkaTimeOut
	if duration != 0 {
		sendConfig.Producer.Timeout = duration
	}
	return &KafkaConfig{
		addressList: address,
		topic:       topic,
		config:      sendConfig,
	}
}
