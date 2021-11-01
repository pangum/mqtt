package mqtt

import (
	`github.com/pangum/pangu-mqtt/core`
)

type will struct {
	// 是否开启
	Enabled bool `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 主题
	Topic string `json:"topic" yaml:"topic" xml:"topic" toml:"topic"`
	// 透传数据
	Payload string `json:"payload" yaml:"payload" xml:"payload" toml:"payload"`
	// 服务等级
	Qos core.Qos `json:"qos" yaml:"qos" xml:"qos" toml:"qos"`
	// 是否保留消息
	Retained bool `json:"retained" yaml:"retained" xml:"retained" toml:"retained"`
}
