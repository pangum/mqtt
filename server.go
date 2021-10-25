package mqtt

type server struct {
	// 标签，后续
	Label string `json:"label" yaml:"label" xml:"label" toml:"label" validate:"required"`
	// 地址
	Broker broker `json:"broker" yaml:"broker" xml:"broker" toml:"broker" validate:"required_without=required,dive"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
