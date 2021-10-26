package mqtt

type subscription struct {
	topic   string
	handler handler
	options []subscribeOption
}
