package mqtt

type config struct {
	// 地址
	Broker broker `json:"broker" yaml:"broker" xml:"broker" toml:"broker" validate:"required_without=Servers"`
	// 服务器列表
	Servers []server `json:"servers" yaml:"servers" xml:"servers" toml:"servers" validate:"required_without=Broker,dive"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
