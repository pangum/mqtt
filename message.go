package mqtt

import (
	mqtt `github.com/eclipse/paho.mqtt.golang`
)

// Message 消息简单封装
type Message struct {
	mqtt.Message
}