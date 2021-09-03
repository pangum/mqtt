package verify

// AclRequest 请求
type AclRequest struct {
	Request

	// 挂载点
	MountPoint string `json:"mountPoint,omitempty" param:"mountPoint" query:"mountPoint"`
	// 主题
	Topic string `json:"topic" param:"topic" query:"topic"`
	// 操作类型
	Type AclType `json:"type,omitempty,string" param:"type" query:"type"`
}
