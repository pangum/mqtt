package mqtt

var (
	_            = DisableDefault
	_            = DisableDefaults
	_ fillOption = (*optionDefaults)(nil)
)

type optionDefaults struct {
	defaults bool
}

// DisableDefault 是否处理默认值
func DisableDefault() *optionDefaults {
	return DisableDefaults()
}

// DisableDefaults 是否处理默认值
func DisableDefaults() *optionDefaults {
	return &optionDefaults{
		defaults: false,
	}
}

func (d *optionDefaults) applyFill(options *fillOptions) {
	options.defaults = d.defaults
}
