package mqtt

var (
	_                 = DisableValidate
	_                 = DisableValidates
	_ publishOption   = (*optionValidate)(nil)
	_ subscribeOption = (*optionValidate)(nil)
	_ messageOption   = (*optionValidate)(nil)
)

type optionValidate struct {
	validates bool
}

// DisableValidate 是否验证数据
func DisableValidate() *optionValidate {
	return DisableValidates()
}

// DisableValidates 是否验证数据
func DisableValidates() *optionValidate {
	return &optionValidate{
		validates: false,
	}
}

func (v *optionValidate) applyPublish(options *publishOptions) {
	options.validates = v.validates
}

func (v *optionValidate) applySubscribe(options *subscribeOptions) {
	options.validates = v.validates
}

func (v *optionValidate) applyMessage(options *messageOptions) {
	options.validates = v.validates
}
