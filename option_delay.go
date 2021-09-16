package mqtt

import (
	`time`
)

var _ publishOption = (*optionDelay)(nil)

type optionDelay struct {
	delay time.Duration
}

// Delay 配置延时发送
func Delay(delay time.Duration) *optionDelay {
	return &optionDelay{
		delay: delay,
	}
}

func (d *optionDelay) applyPublish(options *publishOptions) {
	options.delay = d.delay
}
