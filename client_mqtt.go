package mqtt

import (
	`fmt`
	`time`

	`github.com/eclipse/paho.mqtt.golang`
	`github.com/goexl/gox/field`
)

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
	_messageOptions := c.getMessageOptions(_options.label)
	if payload, err = _messageOptions.marshal(payload); nil != err {
		return
	}

	// 使用MQTT内置的延迟功能实现延迟发送
	if 0 != _options.delay {
		if _options.delay <= c.delayMin {
			_options.delay = c.delayMin
		} else if _options.delay >= c.delayMax {
			_options.delay = c.delayMax
		}
		topic = fmt.Sprintf(`$delayed/%d/%s`, _options.delay/time.Second, topic)
	}
	c.logger.Debug(
		`发送消息`,
		field.String(`topic`, topic),
		field.Int(`size`, len(payload.([]byte))),
		field.Bool(`retained`, _options.retained),
		field.Int8(`qos`, int8(_options.qos)),
	)
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

	_messageOptions := c.getMessageOptions(_options.label)
	token := client.Subscribe(topic, byte(_options.qos), func(client mqtt.Client, message mqtt.Message) {
		go c.consume(handler, client, message, _messageOptions)
	})
	go func() {
		token.Wait()
	}()

	// 存储订阅关系
	if !_options.save {
		return
	}
	c.subscriptions = append(c.subscriptions, subscription{
		topic:   topic,
		handler: handler,
		options: opts,
	})

	return
}

func (c *Client) Resubscribe() (successes []string, fails []string, err error) {
	successes = make([]string, 0)
	fails = make([]string, 0)

	for _, _subscription := range c.subscriptions {
		_subscription.options = append(_subscription.options, NotSave())
		if err = c.Subscribe(_subscription.topic, _subscription.handler, _subscription.options...); nil != err {
			fails = append(fails, _subscription.topic)
		} else {
			successes = append(successes, _subscription.topic)
		}
	}

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

func (c *Client) consume(handler handler, _ mqtt.Client, original mqtt.Message, options *messageOptions) {
	message := newMessage(original, options)
	if err := handler.OnMessage(message); nil == err {
		original.Ack()
	}
}
