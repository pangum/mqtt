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
	Keepalive time.Duration `default:"20s" json:"keepalive" yaml:"keepalive" xml:"keepalive" toml:"keepalive"`
	// 客户端编号
	ClientId string `json:"clientId" yaml:"clientId" xml:"clientId" toml:"clientId"`
	// 是否有序
	Order bool `json:"order" yaml:"order" xml:"order" toml:"order"`
	// 重连时是保留参数
	Will will `json:"will" yaml:"will" xml:"will" toml:"will"`
}
