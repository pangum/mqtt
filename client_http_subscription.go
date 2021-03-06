package mqtt

func (c *Client) Subscription(opts ...brokersOption) []string {
	_options := defaultBrokersOptions()
	for _, opt := range opts {
		opt.applyBrokers(_options)
	}

	return c.brokerCache[_options.label].urls()
}
