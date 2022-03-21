package mqtt

import (
	`github.com/pangum/mqtt/core`
)

type (
	subscribeOption interface {
		applySubscribe(options *subscribeOptions)
	}

	subscribeOptions struct {
		*options
		*messageOptions

		qos  core.Qos
		save bool
	}
)

func defaultSubscribeOptions() *subscribeOptions {
	return &subscribeOptions{
		options:        defaultOptions(),
		messageOptions: defaultMessageOptions(),

		qos:  core.Qos1,
		save: true,
	}
}
