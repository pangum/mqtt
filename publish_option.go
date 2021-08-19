package mqtt

type publishOption interface {
	applyPublish(options *publishOptions)
}
