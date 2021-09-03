package mqtt

const (
	serializerUnknown serializer = ""
	serializerJson    serializer = "json"
	serializerProto   serializer = "proto"
	serializerMsgpack serializer = "msgpack"
	serializerXml     serializer = "xml"
	serializerString  serializer = "string"
	serializerBytes   serializer = "bytes"
)

type serializer string
