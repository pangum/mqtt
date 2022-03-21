package mqtt

import (
	`github.com/eclipse/paho.mqtt.golang`
	`github.com/goexl/gox`
)

// Message 消息封装
type Message struct {
	original mqtt.Message
	options  *messageOptions

	_ gox.CannotCopy
}

func newMessage(original mqtt.Message, options *messageOptions) *Message {
	return &Message{
		original: original,
		options:  options,
	}
}

func (m *Message) Fill(value interface{}, opts ...messageOption) (err error) {
	for _, opt := range opts {
		opt.applyMessage(m.options)
	}
	err = m.options.unmarshal(m.original.Payload(), value)

	return
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
