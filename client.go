package mqtt

import (
	`sync`

	`github.com/eclipse/paho.mqtt.golang`
)

// Client MQTT客户端封装
type Client struct {
	clientCache     map[string]mqtt.Client
	optionsCache    map[string]*mqtt.ClientOptions
	brokersCache    map[string][]string
	serializerCache map[string]serializer
	httpCache       map[string]http

	mutex sync.Mutex
}

func (c *Client) Brokers(opts ...brokersOption) []string {
	_options := defaultBrokersOptions()
	for _, opt := range opts {
		opt.applyBrokers(_options)
	}

	return c.brokersCache[_options.label]
}

func (c *Client) getClient(label string) (client mqtt.Client, err error) {
	var exist bool
	if client, exist = c.clientCache[label]; exist {
		return
	}

	client = mqtt.NewClient(c.optionsCache[label])
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	c.clientCache[label] = client

	return
}

func (c *Client) getSerializer(label string, original serializer) (serializer serializer) {
	if serializerUnknown == original {
		serializer = c.serializerCache[label]
	} else {
		serializer = original
	}

	return
}
