package mqtt

import (
	`time`
)

type mqttOptions struct {
	// 用户名
	Username string `json:"username" yaml:"username" xml:"username"`
	// 密码
	Password string `json:"password" yaml:"password" xml:"password"`
	// 保活时间
	Keepalive time.Duration `default:"20s" json:"keepalive" yaml:"keepalive" xml:"keepalive"`
}
