package mqtt

import (
	`encoding/json`
	`encoding/xml`

	`github.com/eclipse/paho.mqtt.golang`
	`github.com/vmihailenco/msgpack/v5`
	`google.golang.org/protobuf/proto`
)

// Client MQTT客户端封装
type Client struct {
	clientCache  map[string]mqtt.Client
	optionsCache map[string]*mqtt.ClientOptions
}

func (c *Client) Init(opts ...initOption) {

}

func (c *Client) Publish(topic string, payload interface{}, opts ...publishOption) (err error) {
	options := defaultPublishOptions()
	for _, opt := range opts {
		opt.applyPublish(options)
	}

	var client mqtt.Client
	if client, err = c.getClient(options.options); nil != err {
		return
	}

	// 序列化数据
	var bytes []byte
	switch options.format {
	case "proto":
		bytes, err = proto.Marshal(payload.(proto.Message))
	case "json":
		bytes, err = json.Marshal(payload)
	case "xml":
		bytes, err = xml.Marshal(payload)
	case "msgpack":
		bytes, err = msgpack.Marshal(payload)
	}
	if nil != err {
		return
	}

	token := client.Publish(topic, options.qos, options.retained, bytes)
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
	if client, err = c.getClient(options.options); nil != err {
		return
	}

	token := client.Subscribe(topic, options.qos, func(client mqtt.Client, message mqtt.Message) {
		handler.OnMessage(&Message{Message: message})
	})
	go func() {
		<-token.Done()
	}()

	return
}

func (c *Client) getClient(options *options) (client mqtt.Client, err error) {
	var exist bool
	if client, exist = c.clientCache[options.label]; exist {
		return
	}

	client = mqtt.NewClient(c.optionsCache[options.label])

	return
}
