package mqtt

type panguConfig struct {
	// MQTT的配置信息
	Mqtt config `json:"mqtt" yaml:"mqtt" xml:"mqtt" toml:"mqtt" validate:"required"`
}
