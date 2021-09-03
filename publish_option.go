package mqtt

import (
	`github.com/storezhang/pangu-mqtt/core`
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
