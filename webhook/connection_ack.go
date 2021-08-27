package webhook

// ConnectionAck 服务端准备下发连接应答报文时
type ConnectionAck struct {
	Connection

	// "success" 表示成功，其它表示失败的原因
	ConnAck string `json:"conn_ack"`
}
