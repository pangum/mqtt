package mqtt

import (
	`time`
)

type mqttOptions struct {
	// 用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
	// 序列化器，默认使用Msgpack做序列化
	Serializer serializer `default:"msgpack" json:"serializer" yaml:"serializer" xml:"serializer" toml:"serializer" validate:"oneof=json msgpack proto xml"`
	// 保活时间
	Keepalive time.Duration `default:"60s" json:"keepalive" yaml:"keepalive" xml:"keepalive" toml:"keepalive"`
	// 客户端编号
	Clientid string `json:"clientid" yaml:"clientid" xml:"clientid" toml:"clientid"`
	// 重连
	Reconnect reconnect `json:"reconnect" yaml:"reconnect" xml:"reconnect" toml:"reconnect"`
	// 会话
	Session session `json:"session" yaml:"session" xml:"session" toml:"session"`
	// 重试
	Retry retry `json:"retry" yaml:"retry" xml:"retry" toml:"retry"`
	// 超时
	Timeout timeout `json:"timeout" yaml:"timeout" xml:"timeout" toml:"timeout"`
	// 是否有序
	Order bool `json:"order" yaml:"order" xml:"order" toml:"order"`
	// 重连时是保留参数
	Will will `json:"will" yaml:"will" xml:"will" toml:"will"`
}
