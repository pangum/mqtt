package mqtt

import (
	`fmt`
)

type broker struct {
	// 域名
	Domain string `json:"domain" yaml:"domain" xml:"domain" toml:"domain" validate:"required,hostname"`

	// Websocket端口
	Ws int `json:"ws" yaml:"ws" xml:"ws" toml:"ws" validate:"required_without_all=Wss Mqtt Mqtts"`
	// Websocket端口
	Wss int `json:"wss" yaml:"wss" xml:"wss" toml:"wss" validate:"required_without_all=Ws Mqtt Mqtts"`
	// 连接路径
	Path string `json:"path" yaml:"path" xml:"path" toml:"path"`
	// Mqtt端口
	Mqtt int `json:"mqtt" yaml:"mqtt" xml:"mqtt" toml:"mqtt" validate:"required_without_all=Ws Wss Mqtts"`
	// Mqtt SSL端口
	Mqtts int `json:"mqtts" yaml:"mqtts" xml:"mqtts" toml:"mqtts" validate:"required_without_all=Ws Wss Mqtt"`
	// Http接口
	Http http `json:"http" yaml:"http" xml:"http" toml:"http" validate:"structonly"`
}

func (b *broker) validate() bool {
	return `` != b.Domain
}

func (b broker) urls() (urls []string) {
	urls = make([]string, 0)
	if 0 != b.Mqtts {
		urls = append(urls, fmt.Sprintf(connectionFormatter, protocolMqtts, b.Domain, b.Mqtts))
	}
	if 0 != b.Mqtt {
		urls = append(urls, fmt.Sprintf(connectionFormatter, protocolMqtt, b.Domain, b.Mqtt))
	}
	if 0 != b.Wss {
		urls = append(urls, fmt.Sprintf(`%s/%s`, fmt.Sprintf(connectionFormatter, protocolWss, b.Domain, b.Wss), b.Path))
	}
	if 0 != b.Ws {
		urls = append(urls, fmt.Sprintf(`%s/%s`, fmt.Sprintf(connectionFormatter, protocolWss, b.Domain, b.Ws), b.Path))
	}

	return
}

func (b *broker) best() (addr string) {
	var _protocol protocol
	var port int
	defer func() {
		addr = fmt.Sprintf(connectionFormatter, _protocol, b.Domain, port)
	}()

	_protocol = protocolMqtts
	port = b.Mqtts
	if 0 != port {
		return
	}

	_protocol = protocolMqtt
	port = b.Mqtt
	if 0 != port {
		return
	}

	_protocol = protocolWss
	port = b.Wss
	if 0 != port {
		return
	}

	_protocol = protocolWs
	port = b.Ws
	if 0 != port {
		return
	}

	return
}
