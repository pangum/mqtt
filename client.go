package mqtt

import (
	`encoding/json`
	`encoding/xml`
	`sync`
	`time`

	`github.com/eclipse/paho.mqtt.golang`
	`github.com/vmihailenco/msgpack/v5`
	`google.golang.org/protobuf/proto`
)

// Client MQTT客户端封装
type Client struct {
	clientCache     map[string]mqtt.Client
	optionsCache    map[string]*mqtt.ClientOptions
	brokersCache    map[string][]string
	serializerCache map[string]serializer

	mutex sync.Mutex
}

func (c *Client) Brokers(opts ...brokersOption) []string {
	_options := defaultBrokersOptions()
	for _, opt := range opts {
		opt.applyBrokers(_options)
	}

	return c.brokersCache[_options.label]
}

func (c *Client) Publish(topic string, payload interface{}, opts ...publishOption) (err error) {
	_options := defaultPublishOptions()
	for _, opt := range opts {
		opt.applyPublish(_options)
	}

	var client mqtt.Client
	if client, err = c.getClient(_options.options.label); nil != err {
		return
	}

	// 序列化数据
	_serializer := c.getSerializer(_options.label, _options.serializer)
	switch _serializer {
	case serializerProto:
		payload, err = proto.Marshal(payload.(proto.Message))
	case serializerJson:
		payload, err = json.Marshal(payload)
	case serializerXml:
		payload, err = xml.Marshal(payload)
	case serializerMsgpack:
		payload, err = msgpack.Marshal(payload)
	case serializerBytes:
		payload = payload.([]byte)
	case serializerString:
		payload = payload.(string)
	}
	if nil != err {
		return
	}

	token := client.Publish(topic, byte(_options.qos), _options.retained, payload)
	go func() {
		<-token.Done()
	}()

	return
}

func (c *Client) Subscribe(topic string, handler handler, opts ...subscribeOption) (err error) {
	_options := defaultSubscribeOptions()
	for _, opt := range opts {
		opt.applySubscribe(_options)
	}

	var client mqtt.Client
	if client, err = c.getClient(_options.options.label); nil != err {
		return
	}

	token := client.Subscribe(topic, byte(_options.qos), func(client mqtt.Client, message mqtt.Message) {
		go c.consume(handler, client, message)
	})
	go func() {
		token.Wait()
	}()

	return
}

func (c *Client) Disconnect(duration time.Duration, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	var client mqtt.Client
	if client, err = c.getClient(_options.label); nil != err {
		return
	}
	client.Disconnect(uint(duration / 1000))

	return
}

func (c *Client) consume(handler handler, _ mqtt.Client, message mqtt.Message) {
	if err := handler.OnMessage(&Message{Message: message}); nil == err {
		message.Ack()
	}
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
