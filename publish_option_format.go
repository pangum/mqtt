package mqtt

var _ publishOption = (*publishOptionFormat)(nil)

type publishOptionFormat struct {
	format format
}

// Proto 谷歌Protocol Buffer序列化
func Proto() *publishOptionFormat {
	return &publishOptionFormat{
		format: formatProto,
	}
}

// JSON 使用JSON序列化
func JSON() *publishOptionFormat {
	return &publishOptionFormat{
		format: formatJson,
	}
}

// XML 使用XML序列化
func XML() *publishOptionFormat {
	return &publishOptionFormat{
		format: formatXml,
	}
}

// Msgpack 使用Msgpack序列化
func Msgpack() *publishOptionFormat {
	return &publishOptionFormat{
		format: formatMsgpack,
	}
}

// Bytes 原始数据
func Bytes() *publishOptionFormat {
	return &publishOptionFormat{
		format: formatBytes,
	}
}

// String 字符串数据
func String() *publishOptionFormat {
	return &publishOptionFormat{
		format: formatString,
	}
}

func (f *publishOptionFormat) applyPublish(options *publishOptions) {
	options.format = f.format
}
