package mqtt

import (
	`github.com/storezhang/pangu-mqtt/core`
)

type (
	subscribeOption interface {
		applySubscribe(options *subscribeOptions)
	}

	subscribeOptions struct {
		*options

		format string
		qos    core.Qos
	}
)

func defaultSubscribeOptions() *subscribeOptions {
	return &subscribeOptions{
		options: defaultOptions(),

		format: "proto",
		qos:    core.Qos1,
	}
}
