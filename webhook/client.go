package webhook

// Connection 所有连接事件的基类
type Connection struct {
	Request

	// 客户端源地址
	IpAddress string `json:"ipaddress"`
	// 客户端申请的心跳保活时间
	Keepalive int32 `json:"keepalive"`
	// 协议版本号
	ProtoVersion int32 `json:"proto_ver"`
}
