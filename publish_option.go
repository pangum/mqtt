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

func defaultPublishOptions() *publishOptions {
	return &publishOptions{
		options: defaultOptions(),

		serializer: serializerUnknown,
		qos:        core.Qos1,
		retained:   false,
	}
}
