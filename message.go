package mqtt

import (
	`github.com/eclipse/paho.mqtt.golang`
	`github.com/goexl/gox`
	`github.com/goexl/mengpo`
	`github.com/goexl/xiren`
)

// Message 消息封装
type Message struct {
	original   mqtt.Message
	serializer serializer

	_ gox.CannotCopy
}

func (m *Message) Fill(value interface{}, opts ...fillOption) (err error) {
	_options := defaultFillOptions()
	for _, opt := range opts {
		opt.applyFill(_options)
	}

	if err = m.serializer.Unmarshal(m.original.Payload(), value); nil != err {
		return
	}

	// 加载默认值
	if _options.defaults {
		err = mengpo.Set(value)
	}
	if nil != err {
		return
	}

	// 数据验证
	if _options.validates {
		err = xiren.Struct(value)
	}

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
