package mqtt

type http struct {
	// 地址
	Url string `json:"url" yaml:"url" xml:"url" toml:"url" validate:"required"`
	// 端点
	Endpoint string `default:"/api/v4" json:"endpoint" yaml:"endpoint" xml:"endpoint" toml:"endpoint"`
	// 用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
}
