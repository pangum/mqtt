package webhook

// Disconnected 客户端连接层在准备关闭时
type Disconnected struct {
	Request

	// 错误原因
	Reason string `json:"reason"`
}
