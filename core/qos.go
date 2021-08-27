package core

const (
	// Qos0 消息可能会被丢掉
	Qos0 Qos = 0
	// Qos1 消息至少送达一次（At the least once delivery）
	Qos1 Qos = 1
	// Qos2 消息只送达一次（Exactly once delivery）
	Qos2 Qos = 2
)

// Qos 服务等级
type Qos byte
