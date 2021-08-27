package core

// VerifyRequest 请求
type VerifyRequest struct {
	// 用户名
	Username string `json:"username" param:"username" query:"username"`
	// 客户端编号
	ClientId string `json:"clientId" param:"clientId" query:"clientId"`
	// 客户端IP地址
	Ip string `json:"ip" param:"ip" query:"ip"`
	// 客户端接入协议
	Protocol string `json:"protocol" param:"protocol" query:"protocol"`
}
