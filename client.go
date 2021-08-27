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
	clientCache  map[string]mqtt.Client
	optionsCache map[string]*mqtt.ClientOptions
	brokersCache map[string][]string

	mutex sync.Mutex
}

func (c *Client) Brokers(opts ...brokersOption) []string {
	options := defaultBrokersOptions()
	for _, opt := range opts {
		opt.applyBrokers(options)
	}

	return c.brokersCache[options.label]
}

func (c *Client) Publish(topic string, payload interface{}, opts ...publishOption) (err error) {
	options := defaultPublishOptions()
	for _, opt := range opts {
		opt.applyPublish(options)
	}

	var client mqtt.Client
	if client, err = c.getClient(options.options.label); nil != err {
		return
	}

	// 序列化数据
	switch options.format {
	case formatProto:
		payload, err = proto.Marshal(payload.(proto.Message))
	case formatJson:
		payload, err = json.Marshal(payload)
	case formatXml:
		payload, err = xml.Marshal(payload)
	case formatMsgpack:
		payload, err = msgpack.Marshal(payload)
	case formatBytes:
		payload = payload.([]byte)
	case formatString:
		payload = payload.(string)
	}
	if nil != err {
		return
	}

	token := client.Publish(topic, byte(options.qos), options.retained, payload)
	go func() {
		<-token.Done()
	}()

	return
}

func (c *Client) Subscribe(topic string, handler handler, opts ...subscribeOption) (err error) {
	options := defaultSubscribeOptions()
	for _, opt := range opts {
		opt.applySubscribe(options)
	}

	var client mqtt.Client
	if client, err = c.getClient(options.options.label); nil != err {
		return
	}

	token := client.Subscribe(topic, byte(options.qos), func(client mqtt.Client, message mqtt.Message) {
		go c.consume(handler, client, message)
	})
	go func() {
		token.Wait()
	}()

	return
}

func (c *Client) Disconnect(duration time.Duration, opts ...option) (err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	var client mqtt.Client
	if client, err = c.getClient(options.label); nil != err {
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
