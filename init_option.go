package mqtt

type initOption interface {
	applyInit()
}
