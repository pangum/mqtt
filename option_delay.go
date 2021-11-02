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

// Time 配置某个时刻执行
func Time(t time.Time) *optionDelay {
	return Until(t)
}

// Timestamp 配置某个时刻执行
func Timestamp(timestamp time.Time) *optionDelay {
	return Until(timestamp)
}

// Until 配置直到某个时间点触发
func Until(until time.Time) *optionDelay {
	return &optionDelay{
		delay: until.Sub(time.Now()),
	}
}

func (d *optionDelay) applyPublish(options *publishOptions) {
	options.delay = d.delay
}
