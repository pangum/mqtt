package mqtt

type (
	fillOption interface {
		applyFill(options *fillOptions)
	}

	fillOptions struct {
		defaults  bool
		validates bool
	}
)

func defaultFillOptions() *fillOptions {
	return &fillOptions{
		defaults:  true,
		validates: true,
	}
}
