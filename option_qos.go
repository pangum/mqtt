package mqtt

import (
	`github.com/storezhang/pangu-mqtt/core`
)

var (
	_ publishOption   = (*optionQos)(nil)
	_ subscribeOption = (*optionQos)(nil)
)

type optionQos struct {
	qos core.Qos
}

// Qos0 0等级消息，MQTT不保证消息到达
func Qos0() *optionQos {
	return &optionQos{
		qos: core.Qos0,
	}
}

// Qos1 1等级消息，MQTT保证消息至少会被消费一次
func Qos1() *optionQos {
	return &optionQos{
		qos: core.Qos1,
	}
}

// Qos2 2等级消息，MQTT保证消息有且只能被消费一次，此等级最高，会消耗性能，建议不要大量使用
func Qos2() *optionQos {
	return &optionQos{
		qos: core.Qos2,
	}
}

func (q *optionQos) applyPublish(options *publishOptions) {
	options.qos = q.qos
}

func (q *optionQos) applySubscribe(options *subscribeOptions) {
	options.qos = q.qos
}
