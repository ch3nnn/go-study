package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

// AsyncKafkaProducer kafka异步生产者
type AsyncKafkaProducer struct {
	KafkaConfig *KafkaConfig         // 共有的kafka生产者配置在这个里面
	producer    sarama.AsyncProducer // 异步生产者
	isClose     chan struct{}        // 监听producer是否可以关闭
}

// NewKafkaProducer 创建kafka异步生产者实例，并初始化参数
func (k *AsyncKafkaProducer) NewKafkaProducer(address []string, topic string, duration time.Duration) {
	if len(address) == 0 {
		panic(addressError)
	}
	// 配置异步producer启动参数
	k.KafkaConfig = NewProducerByMessage(address, topic, duration)
	k.isClose = make(chan struct{}, 2)
	// 启动kafka异步producer
	k.Run()
}

// Send kafka异步发送消息
func (k *AsyncKafkaProducer) Send(value []byte) {
	// 如果实例或者配置为空，直接返回。如果发送数据为空也直接返回
	if k == nil || k.KafkaConfig == nil || value == nil {
		return
	}
	// 封装消息实例
	msg := &sarama.ProducerMessage{
		Topic: k.KafkaConfig.topic,
		Value: sarama.ByteEncoder(value),
	}
	// 这里一般不会出现，producer实例为空时，表示创建异步producer失败
	if k.producer == nil {
		k.Run()
	}
	select {
	// producer 出现error需要重新启动
	case <-k.isClose:
		if k.producer != nil {
			// 收到可以关闭producer的消息isClose,关闭producer并重启
			k.producer.Close()
			k.Run()
		}
		// 直接返回，此条消息浪费掉了，如果后期需要收集未发送成功的消息可以在此收集，输出到日志或者等待producer重启成功后再发送
		return
	default:
		// 正常情况发送消息
		k.producer.Input() <- msg
	}
}

// Run kafka异步生产者
func (k *AsyncKafkaProducer) Run() {
	if k == nil || k.KafkaConfig == nil {
		return
	}
	// 创建异步producer
	producer, err := sarama.NewAsyncProducer(k.KafkaConfig.addressList, k.KafkaConfig.config)
	// 如果创建失败主动置空k.producer，否则producer不为空，在重启的时候k.producer是会有值的
	if err != nil {
		k.isClose <- struct{}{}
		k.producer = nil
		panic(newAsyncProducerError)
	}
	if producer == nil {
		k.isClose <- struct{}{}
		k.producer = nil
		panic(asyncProducerNullError)
	}
	// 如果创建成功为实例 k 的 producer 赋值
	k.producer = producer
	go func(p sarama.AsyncProducer) {
		producerError := p.Errors()
		producerMessage := p.Successes()
		for {
			select {
			// 出现了错误
			case rc := <-producerError:
				if rc != nil {
					// 标记producer出现error，在send时会监听到这个标记
					k.isClose <- struct{}{}
					panic(sendMessageError)
				}
				return
			case res := <-producerMessage:
				// 发送成功
				res.Value.Encode()
			}
		}
	}(producer)

}
