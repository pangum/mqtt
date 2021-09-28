package mqtt

type session struct {
	// 清除会话
	Clean bool `default:"false" json:"clean" yaml:"clean" xml:"clean" toml:"clean"`
}
