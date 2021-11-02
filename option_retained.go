package mqtt

var _ publishOption = (*optionRetained)(nil)

type optionRetained struct{}

// Retained 配置是否保存到服务器
func Retained() *optionRetained {
	return &optionRetained{}
}

func (r *optionRetained) applyPublish(options *publishOptions) {
	options.retained = true
}
