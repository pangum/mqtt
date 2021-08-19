package mqtt

var _ publishOption = (*publishOptionFormat)(nil)

type publishOptionFormat struct {
	format string
}

// Proto 谷歌Protocol Buffer序列化
func Proto() *publishOptionFormat {
	return &publishOptionFormat{
		format: "proto",
	}
}

// JSON 使用JSON序列化
func JSON() *publishOptionFormat {
	return &publishOptionFormat{
		format: "json",
	}
}

// XML 使用XML序列化
func XML() *publishOptionFormat {
	return &publishOptionFormat{
		format: "xml",
	}
}

// Msgpack 使用Msgpack序列化
func Msgpack() *publishOptionFormat {
	return &publishOptionFormat{
		format: "msgpack",
	}
}

func (q *publishOptionFormat) applyPublish(options *publishOptions) {
	options.format = q.format
}
