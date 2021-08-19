package mqtt

type option interface {
	apply(options *options)
}
