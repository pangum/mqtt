package mqtt

type config struct {
	// 地址列表
	Brokers []string `json:"brokers" yaml:"brokers" xml:"brokers" toml:"brokers" validate:"required_without=Servers,dive,url"`
	// 服务器列表
	Servers []server `json:"servers" yaml:"servers" xml:"servers" toml:"servers" validate:"required_without=Brokers,dive"`
	// Http接口
	Http http `json:"http" yaml:"http" xml:"http" toml:"http" validate:"structonly"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
