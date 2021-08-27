package core

const (
	// AclTypeUnknown 未知
	AclTypeUnknown AclType = 0
	// AclTypeSubscribe 订阅
	AclTypeSubscribe AclType = 1
	// AclTypePublish 发布
	AclTypePublish AclType = 2
)

// AclType 权限类型
type AclType uint8
