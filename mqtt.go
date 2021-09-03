package mqtt

import (
	`github.com/eclipse/paho.mqtt.golang`
	`github.com/storezhang/pangu`
)

func newMqtt(config *pangu.Config) (client *Client, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}

	mqttConfig := panguConfig.Mqtt

	// 加载默认连接
	brokersCache := make(map[string][]string)
	optionsCache := make(map[string]*mqtt.ClientOptions)
	serializerCache := make(map[string]serializer)
	if 0 != len(mqttConfig.Brokers) {
		_defaultOptions := mqtt.NewClientOptions()
		for _, broker := range mqttConfig.Brokers {
			_defaultOptions.AddBroker(broker)
		}
		_defaultOptions.SetUsername(mqttConfig.Options.Username)
		_defaultOptions.SetPassword(mqttConfig.Options.Password)
		_defaultOptions.SetKeepAlive(mqttConfig.Options.Keepalive)
		_defaultOptions.SetClientID(mqttConfig.Options.ClientId)
		if mqttConfig.Options.Will.Enabled {
			_defaultOptions.SetWill(
				mqttConfig.Options.Will.Topic,
				mqttConfig.Options.Will.Payload,
				byte(mqttConfig.Options.Will.Qos),
				mqttConfig.Options.Will.Retained,
			)
		}

		optionsCache[defaultLabel] = _defaultOptions
		brokersCache[defaultLabel] = mqttConfig.Brokers
		serializerCache[defaultLabel] = mqttConfig.Options.Serializer
	}

	// 加载带标签的服务器
	for _, _server := range mqttConfig.Servers {
		serverOptions := mqtt.NewClientOptions()
		for _, broker := range _server.Brokers {
			serverOptions.AddBroker(broker)
		}

		setString(serverOptions.SetUsername, _server.Options.Username, mqttConfig.Options.Username)
		setString(serverOptions.SetPassword, _server.Options.Password, mqttConfig.Options.Password)
		setDuration(serverOptions.SetKeepAlive, _server.Options.Keepalive, mqttConfig.Options.Keepalive)
		setString(serverOptions.SetClientID, _server.Options.ClientId, mqttConfig.Options.ClientId)
		if _server.Options.Will.Enabled {
			serverOptions.SetWill(
				_server.Options.Will.Topic,
				_server.Options.Will.Payload,
				byte(_server.Options.Will.Qos),
				_server.Options.Will.Retained,
			)
		}

		optionsCache[_server.Label] = serverOptions
		brokersCache[_server.Label] = _server.Brokers
		serializerCache[_server.Label] = _server.Options.Serializer
	}

	client = &Client{
		clientCache:     make(map[string]mqtt.Client),
		optionsCache:    optionsCache,
		brokersCache:    brokersCache,
		serializerCache: serializerCache,
	}

	return
}
