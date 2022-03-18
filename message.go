package mqtt

import (
	`github.com/eclipse/paho.mqtt.golang`
	`github.com/goexl/gox`
)

// Message 消息封装
type Message struct {
	original   mqtt.Message
	serializer serializer

	_ gox.CannotCopy
}

func (m *Message) Fill(value interface{}) error {
	return m.serializer.Unmarshal(m.original.Payload(), value)
}

func (m *Message) Duplicate() bool {
	return m.original.Duplicate()
}

func (m *Message) Qos() byte {
	return m.original.Qos()
}

func (m *Message) Retained() bool {
	return m.original.Retained()
}

func (m *Message) Topic() string {
	return m.original.Topic()
}

func (m *Message) MessageId() uint16 {
	return m.original.MessageID()
}
