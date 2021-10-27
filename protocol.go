package mqtt

const (
	protocolMqtts protocol = "mqtts"
	protocolMqtt  protocol = "mqtt"
	protocolWss   protocol = "wss"
	protocolWs    protocol = "ws"
)

type protocol string
