package mqtt

import (
	`time`
)

type reconnect struct {
	// 自动重连
	Auto bool `default:"true" json:"auto" yaml:"auto" xml:"auto" toml:"auto"`
	// 重连间隔
	Interval time.Duration `default:"10s" json:"interval" yaml:"interval" xml:"interval" toml:"interval"`
}
