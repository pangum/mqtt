package webhook

// Unsubscribe 完成取消订阅操作后
type Unsubscribe struct {
	Request

	// 将订阅的主题
	Topic string `json:"topic"`
}
