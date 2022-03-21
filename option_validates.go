package mqtt

var (
	_            = DisableValidate
	_            = DisableValidates
	_ fillOption = (*optionValidate)(nil)
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

func (v *optionValidate) applyFill(options *fillOptions) {
	options.validates = v.validates
}
