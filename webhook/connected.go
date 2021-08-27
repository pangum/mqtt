package webhook

// Connected 客户端认证完成并成功接入系统后
type Connected struct {
	Connection

	// 时间戳（秒）
	ConnectedAt int32 `json:"connected_at"`
}
