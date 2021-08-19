package mqtt

type subscribeOption interface {
	applySubscribe(options *subscribeOptions)
}
