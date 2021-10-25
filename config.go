package mqtt

type config struct {
	// 地址
	Broker broker `json:"broker" yaml:"broker" xml:"broker" toml:"broker" validate:"required_without=Brokers"`
	// 服务器列表
	Brokers []labeledServer `json:"brokers" yaml:"brokers" xml:"brokers" toml:"brokers" validate:"required_without=Broker,dive"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
