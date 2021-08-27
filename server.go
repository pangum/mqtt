package mqtt

type server struct {
	// 标签，后续
	Label string `json:"label" yaml:"label" xml:"label" toml:"label" validate:"required"`
	// 地址
	Broker string `json:"broker" yaml:"broker" xml:"broker" toml:"broker" validate:"required_without=Brokers,url"`
	// 地址列表（集群模式）
	Brokers []string `json:"brokers" yaml:"brokers" xml:"brokers" toml:"brokers" validate:"required_without=Broker,dive,url"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
