package mqtt

type (
	brokersOption interface {
		applyBrokers(options *brokersOptions)
	}

	brokersOptions struct {
		*options
	}
)

func defaultBrokersOptions() *brokersOptions {
	return &brokersOptions{
		options: defaultOptions(),
	}
}
