package webhook

import (
	`github.com/pangum/mqtt/core`
)

// Subscribed 收到订阅报文后，执行client.check_acl鉴权前
type Subscribed struct {
	Request

	// 将订阅的主题
	Topic string `json:"topic"`
	// 订阅参数
	Opts struct {
		// 服务等级
		Qos core.Qos `json:"qos"`
	} `json:"opts"`
}
