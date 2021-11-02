package mqtt

import (
	`time`

	`github.com/pangum/mqtt/core`
)

type (
	publishOption interface {
		applyPublish(options *publishOptions)
	}

	publishOptions struct {
		*options

		serializer serializer
		qos        core.Qos
		retained   bool
		delay      time.Duration
	}
)

// PublishOptions 因为publishOption没有暴露，在外面不方便处理，特意留一个组装各种选项的方法
func PublishOptions(opts ...publishOption) []publishOption {
	return opts
}

func defaultPublishOptions() *publishOptions {
	return &publishOptions{
		options: defaultOptions(),

		serializer: serializerUnknown,
		qos:        core.Qos1,
		retained:   false,
	}
}
