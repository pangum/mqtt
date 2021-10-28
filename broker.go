package mqtt

import (
	`fmt`
)

type broker struct {
	// 域名
	Domain string `json:"domain" yaml:"domain" xml:"domain" toml:"domain" validate:"required,hostname|ip"`

	// Websocket端口
	Ws int `json:"ws" yaml:"ws" xml:"ws" toml:"ws" validate:"omitempty,required_without_all=Wss Mqtt Mqtts"`
	// Websocket端口
	Wss int `json:"wss" yaml:"wss" xml:"wss" toml:"wss" validate:"omitempty,required_without_all=Ws Mqtt Mqtts"`
	// 连接路径
	Path string `default:"mqtt" json:"path" yaml:"path" xml:"path" toml:"path"`
	// Mqtt端口
	Mqtt int `json:"mqtt" yaml:"mqtt" xml:"mqtt" toml:"mqtt" validate:"omitempty,required_without_all=Ws Wss Mqtts"`
	// Mqtt SSL端口
	Mqtts int `json:"mqtts" yaml:"mqtts" xml:"mqtts" toml:"mqtts" validate:"omitempty,required_without_all=Ws Wss Mqtt"`
	// Http接口
	Http http `json:"http" yaml:"http" xml:"http" toml:"http" validate:"structonly"`

	// 连接最佳协议
	Orders []protocol `default:"[mqtts,mqtt,wss,ws]" json:"orders" yaml:"orders" xml:"orders" toml:"orders" validate:"required,dive,oneof=mqtts mqtt wss ws"`
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
	for _, order := range b.Orders {
		switch {
		case protocolMqtts == order && 0 != b.Mqtts:
			addr = fmt.Sprintf(connectionFormatter, protocolMqtts, b.Domain, b.Mqtts)
		case protocolMqtt == order && 0 != b.Mqtt:
			addr = fmt.Sprintf(connectionFormatter, protocolMqtt, b.Domain, b.Mqtt)
		case protocolWss == order && 0 != b.Wss:
			addr = fmt.Sprintf(`%s/%s`, fmt.Sprintf(connectionFormatter, protocolWss, b.Domain, b.Wss), b.Path)
		case protocolWs == order && 0 != b.Ws:
			addr = fmt.Sprintf(`%s/%s`, fmt.Sprintf(connectionFormatter, protocolWs, b.Domain, b.Ws), b.Path)
		}
		if `` != addr {
			break
		}
	}

	return
}
