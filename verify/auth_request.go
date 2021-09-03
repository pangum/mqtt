package verify

// AuthRequest 请求
type AuthRequest struct {
	Request

	// 明文密码
	Password string `json:"password" param:"password" query:"password"`
	// 客户端端口
	Port int32 `json:"port,string" param:"port" query:"port"`
}
