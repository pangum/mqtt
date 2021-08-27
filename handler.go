package mqtt

type handler interface {
	// OnMessage 处理消息
	OnMessage(message *Message) (err error)
}
