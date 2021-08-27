package mqtt

const (
	formatJson    format = "json"
	formatProto   format = "proto"
	formatMsgpack format = "msgpack"
	formatXml     format = "xml"
	formatString  format = "string"
	formatBytes   format = "bytes"
)

type format string
