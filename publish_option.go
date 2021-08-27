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

		format   format
		qos      core.Qos
		retained bool
	}
)

func defaultPublishOptions() *publishOptions {
	return &publishOptions{
		options: defaultOptions(),

		format:   "proto",
		qos:      core.Qos1,
		retained: false,
	}
}
