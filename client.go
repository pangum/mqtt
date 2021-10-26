package mqtt

import (
	`sync`

	`github.com/eclipse/paho.mqtt.golang`
)

// Client MQTT客户端封装
type Client struct {
	clientCache     map[string]mqtt.Client
	optionsCache    map[string]*mqtt.ClientOptions
	brokerCache     map[string]broker
	serializerCache map[string]serializer
	subscriptions   []subscription

	mutex sync.Mutex
}

func newClient(
	optionsCache map[string]*mqtt.ClientOptions,
	brokerCache map[string]broker,
	serializerCache map[string]serializer,
) *Client {
	return &Client{
		clientCache:     make(map[string]mqtt.Client),
		optionsCache:    optionsCache,
		brokerCache:     brokerCache,
		serializerCache: serializerCache,
		subscriptions:   make([]subscription, 0),
	}
}

func (c *Client) Urls(opts ...brokersOption) []string {
	_options := defaultBrokersOptions()
	for _, opt := range opts {
		opt.applyBrokers(_options)
	}

	return c.brokerCache[_options.label].urls()
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
