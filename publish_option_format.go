package mqtt

var _ publishOption = (*publishOptionSerializer)(nil)

type publishOptionSerializer struct {
	serializer serializer
}

// Proto 谷歌Protocol Buffer序列化
func Proto() *publishOptionSerializer {
	return &publishOptionSerializer{
		serializer: serializerProto,
	}
}

// JSON 使用JSON序列化
func JSON() *publishOptionSerializer {
	return &publishOptionSerializer{
		serializer: serializerJson,
	}
}

// XML 使用XML序列化
func XML() *publishOptionSerializer {
	return &publishOptionSerializer{
		serializer: serializerXml,
	}
}

// Msgpack 使用Msgpack序列化
func Msgpack() *publishOptionSerializer {
	return &publishOptionSerializer{
		serializer: serializerMsgpack,
	}
}

// Bytes 原始数据
func Bytes() *publishOptionSerializer {
	return &publishOptionSerializer{
		serializer: serializerBytes,
	}
}

// String 字符串数据
func String() *publishOptionSerializer {
	return &publishOptionSerializer{
		serializer: serializerString,
	}
}

func (f *publishOptionSerializer) applyPublish(options *publishOptions) {
	options.serializer = f.serializer
}
