package mqtt

import (
	`time`
)

type retry struct {
	// 是否开启
	Enable bool `default:"true" json:"enable" yaml:"enable" xml:"enable" toml:"enable"`
	// 间隔
	Interval time.Duration `default:"30s" json:"interval" yaml:"interval" xml:"interval" toml:"interval"`
}
