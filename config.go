package mqtt

type config struct {
	// 地址
	Broker string `json:"broker" yaml:"broker" xml:"broker" validate:"required_without_all=Brokers Servers,url"`
	// 地址列表（集群模式）
	Brokers []string `json:"brokers" yaml:"brokers" xml:"brokers" validate:"required_without_all=Broker Servers,dive,url"`
	// 服务器列表
	Servers []server `json:"servers" yaml:"servers" xml:"servers" validate:"required_without=Broker Brokers,dive"`
	// 选项
	Options mqttOptions `json:"options" yaml:"options" xml:"options"`
}
