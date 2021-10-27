package mqtt

const (
	protocolMqtts protocol = "mqtts"
	protocolMqtt  protocol = "tcp"
	protocolWss   protocol = "wss"
	protocolWs    protocol = "ws"
)

type protocol string
