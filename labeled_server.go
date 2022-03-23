package mqtt

type labeledServer struct {
	// 标签，可以通过标签来选择不同的服务器
	Label string `json:"label" yaml:"label" xml:"label" toml:"label" validate:"required"`
	// 地址
	Broker broker `json:"broker" yaml:"broker" xml:"broker" toml:"broker" validate:"required"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
