package webhook

// Request Webhook回调时的请求参数
type Request struct {
	// 事件名称
	Action Action `json:"action"`
	// 客户端
	Clientid string `json:"clientid"`
	// 客户端用户名，不存在时该值为 "undefined"
	Username string `json:"username"`
}
