package mqtt

type subscribeOptions struct {
	*options

	format string
	qos    byte
}

func defaultSubscribeOptions() *subscribeOptions {
	return &subscribeOptions{
		options: defaultOptions,

		format: "proto",
		qos:    1,
	}
}
