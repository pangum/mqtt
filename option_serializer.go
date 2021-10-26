package mqtt

var _ publishOption = (*optionSerializer)(nil)

type optionSerializer struct {
	serializer serializer
}

// Proto 谷歌Protocol Buffer序列化
func Proto() *optionSerializer {
	return &optionSerializer{
		serializer: serializerProto,
	}
}

// JSON 使用JSON序列化
func JSON() *optionSerializer {
	return &optionSerializer{
		serializer: serializerJson,
	}
}

// XML 使用XML序列化
func XML() *optionSerializer {
	return &optionSerializer{
		serializer: serializerXml,
	}
}

// Msgpack 使用Msgpack序列化
func Msgpack() *optionSerializer {
	return &optionSerializer{
		serializer: serializerMsgpack,
	}
}

// Bytes 原始数据
func Bytes() *optionSerializer {
	return &optionSerializer{
		serializer: serializerBytes,
	}
}

// String 字符串数据
func String() *optionSerializer {
	return &optionSerializer{
		serializer: serializerString,
	}
}

func (f *optionSerializer) applyPublish(options *publishOptions) {
	options.serializer = f.serializer
}
