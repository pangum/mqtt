package webhook

const (
	// ActionUnknown 未知
	ActionUnknown Action = ""
	// ActionClientConnect 处理连接报文，服务端收到客户端的连接报文时
	ActionClientConnect Action = "client_connect"
	// ActionClientConnectionAck 下发连接应答，服务端准备下发连接应答报文时
	ActionClientConnectionAck Action = "client_connack"
	// ActionClientConnected 成功接入，客户端认证完成并成功接入系统后
	ActionClientConnected Action = "client_connected"
	// ActionClientDisconnected 连接断开，客户端连接层在准备关闭时
	ActionClientDisconnected Action = "client_disconnected"
	// ActionSessionTerminated 连接断开，客户端连接层在准备关闭时
	ActionSessionTerminated Action = "session_terminated"
	// ActionClientSubscribe 订阅主题，收到订阅报文后，执行client.check_acl鉴权前
	ActionClientSubscribe Action = "client_subscribe"
	// ActionSessionSubscribe 订阅主题，收到订阅报文后，执行client.check_acl鉴权前
	ActionSessionSubscribe Action = "session_subscribe"
	// ActionClientUnsubscribe 取消订阅，收到取消订阅报文后
	ActionClientUnsubscribe Action = "client_unsubscribe"
	// ActionSessionUnsubscribe 取消订阅，收到取消订阅报文后
	ActionSessionUnsubscribe Action = "session_unsubscribe"
	// ActionMessagePublish 消息发布，服务端在发布（路由）消息前
	ActionMessagePublish Action = "message_publish"
	// ActionMessageDelivered 消息投递，消息准备投递到客户端前
	ActionMessageDelivered Action = "message_delivered"
	// ActionMessageAcked 消息回执，服务端在收到客户端发回的消息ACK后
	ActionMessageAcked Action = "message_acked"
)

// Action MQTT回调事件名称
type Action string
