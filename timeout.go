package mqtt

import (
	`time`
)

type timeout struct {
	// Ping超时
	Ping time.Duration `default:"10s" json:"ping" yaml:"ping" xml:"ping" toml:"ping"`
	// 连接超时
	Connect time.Duration `default:"30s" json:"connect" yaml:"connect" xml:"connect" toml:"connect"`
	// 写入超时
	Write time.Duration `json:"write" yaml:"write" xml:"write" toml:"write"`
}
