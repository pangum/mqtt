package mqtt

import (
	`github.com/goexl/mengpo`
	`github.com/goexl/xiren`
)

type (
	messageOption interface {
		applyMessage(options *messageOptions)
	}

	messageOptions struct {
		serializer serializer
		defaults   bool
		validates  bool
	}
)

func (m *messageOptions) marshal(from interface{}) (to []byte, err error) {
	// 加载默认值
	if m.defaults {
		err = mengpo.Set(from)
	}
	if nil != err {
		return
	}

	// 数据验证
	if m.validates {
		err = xiren.Struct(from)
	}
	if nil != err {
		return
	}

	// 序列化
	to, err = m.serializer.Marshal(from)

	return
}

func (m *messageOptions) unmarshal(bytes []byte, from interface{}) (err error) {
	// 反序列化
	if err = m.serializer.Unmarshal(bytes, from); nil != err {
		return
	}

	// 加载默认值
	if m.defaults {
		err = mengpo.Set(from)
	}
	if nil != err {
		return
	}

	// 数据验证
	if m.validates {
		err = xiren.Struct(from)
	}

	return
}

func newMessageOptions(mqtt mqttOptions) *messageOptions {
	return &messageOptions{
		serializer: mqtt.Serializer,
		defaults:   mqtt.Defaults && mqtt.Default,
		validates:  mqtt.Validates && mqtt.Validate,
	}
}

func defaultMessageOptions() *messageOptions {
	return &messageOptions{
		serializer: serializerJson,
		defaults:   true,
		validates:  true,
	}
}
