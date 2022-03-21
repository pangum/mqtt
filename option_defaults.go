package mqtt

var (
	_                 = DisableDefault
	_                 = DisableDefaults
	_ publishOption   = (*optionDefaults)(nil)
	_ subscribeOption = (*optionDefaults)(nil)
	_ messageOption   = (*optionDefaults)(nil)
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

func (d *optionDefaults) applyPublish(options *publishOptions) {
	options.defaults = d.defaults
}

func (d *optionDefaults) applySubscribe(options *subscribeOptions) {
	options.defaults = d.defaults
}

func (d *optionDefaults) applyMessage(options *messageOptions) {
	options.defaults = d.defaults
}
