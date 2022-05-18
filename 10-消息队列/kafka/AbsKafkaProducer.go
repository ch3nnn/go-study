package kafka

import (
	"time"
)

// AbsKafkaProducer 生产者接口
type AbsKafkaProducer interface {

	//
	// NewKafkaProducer
	//  @Description: 创建生成者
	//  @param address 地址
	//  @param topic 主题
	//  @param duration 超时
	//
	NewKafkaProducer(address []string, topic string, duration time.Duration)

	//
	// Send
	//  @Description: 发送消息
	//  @param value
	//
	Send(value []byte)
}
