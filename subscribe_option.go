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

		format string
		qos    core.Qos
		save   bool
	}
)

func defaultSubscribeOptions() *subscribeOptions {
	return &subscribeOptions{
		options: defaultOptions(),

		format: `proto`,
		qos:    core.Qos1,
		save:   true,
	}
}
