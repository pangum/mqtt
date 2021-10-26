package mqtt

var _ subscribeOption = (*optionSave)(nil)

type optionSave struct {
	save bool
}

// NotSave 配置不存储订阅关系
func NotSave() *optionSave {
	return &optionSave{
		save: false,
	}
}

func (s *optionSave) applySubscribe(options *subscribeOptions) {
	options.save = s.save
}
