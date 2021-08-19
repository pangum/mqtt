package mqtt

type publishOptions struct {
	*options

	format   string
	qos      byte
	retained bool
}

func defaultPublishOptions() *publishOptions {
	return &publishOptions{
		options: defaultOptions,

		format:   "proto",
		qos:      1,
		retained: false,
	}
}
