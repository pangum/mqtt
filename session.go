package mqtt

type session struct {
	// 清除会话
	Clean bool `default:"true" json:"clean" yaml:"clean" xml:"clean" toml:"clean"`
}
