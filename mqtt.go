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
	if 0 != len(mqttConfig.Brokers) {
		defaultOptions := mqtt.NewClientOptions()
		for _, broker := range mqttConfig.Brokers {
			defaultOptions.AddBroker(broker)
		}
		defaultOptions.SetUsername(mqttConfig.Options.Username)
		defaultOptions.SetPassword(mqttConfig.Options.Password)
		defaultOptions.SetKeepAlive(mqttConfig.Options.Keepalive)
		defaultOptions.SetClientID(mqttConfig.Options.ClientId)
		if mqttConfig.Options.Will.Enabled {
			defaultOptions.SetWill(
				mqttConfig.Options.Will.Topic,
				mqttConfig.Options.Will.Payload,
				byte(mqttConfig.Options.Will.Qos),
				mqttConfig.Options.Will.Retained,
			)
		}

		optionsCache[defaultLabel] = defaultOptions
		brokersCache[defaultLabel] = mqttConfig.Brokers
	}

	// 加载带标签的服务器
	for _, server := range mqttConfig.Servers {
		serverOptions := mqtt.NewClientOptions()
		for _, broker := range server.Brokers {
			serverOptions.AddBroker(broker)
		}

		setString(serverOptions.SetUsername, server.Options.Username, mqttConfig.Options.Username)
		setString(serverOptions.SetPassword, server.Options.Password, mqttConfig.Options.Password)
		setDuration(serverOptions.SetKeepAlive, server.Options.Keepalive, mqttConfig.Options.Keepalive)
		setString(serverOptions.SetClientID, server.Options.ClientId, mqttConfig.Options.ClientId)
		if server.Options.Will.Enabled {
			serverOptions.SetWill(
				server.Options.Will.Topic,
				server.Options.Will.Payload,
				byte(server.Options.Will.Qos),
				server.Options.Will.Retained,
			)
		}

		optionsCache[server.Label] = serverOptions
		brokersCache[server.Label] = server.Brokers
	}

	client = &Client{
		clientCache:  make(map[string]mqtt.Client),
		optionsCache: optionsCache,
	}

	return
}
