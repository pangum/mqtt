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
	optionsCache := make(map[string]*mqtt.ClientOptions)
	if "" != mqttConfig.Broker || 0 != len(mqttConfig.Brokers) {
		defaultOptions := mqtt.NewClientOptions()
		if "" != mqttConfig.Broker {
			defaultOptions.AddBroker(mqttConfig.Broker)
		} else {
			for _, broker := range mqttConfig.Brokers {
				defaultOptions.AddBroker(broker)
			}
		}
		defaultOptions.SetUsername(mqttConfig.Options.Username)
		defaultOptions.SetPassword(mqttConfig.Options.Password)
		defaultOptions.SetKeepAlive(mqttConfig.Options.Keepalive)

		optionsCache[defaultLabel] = defaultOptions
	}

	// 加载带标签的服务器
	for _, server := range mqttConfig.Servers {
		serverOptions := mqtt.NewClientOptions()
		if "" != server.Broker {
			serverOptions.AddBroker(mqttConfig.Broker)
		} else {
			for _, broker := range server.Brokers {
				serverOptions.AddBroker(broker)
			}
		}
		if "" != server.Options.Username {
			serverOptions.SetUsername(server.Options.Username)
		} else {
			serverOptions.SetUsername(mqttConfig.Options.Username)
		}
		if "" != server.Options.Password {
			serverOptions.SetPassword(server.Options.Password)
		} else {
			serverOptions.SetPassword(mqttConfig.Options.Password)
		}
		if 0 != server.Options.Keepalive {
			serverOptions.SetKeepAlive(server.Options.Keepalive)
		} else {
			serverOptions.SetKeepAlive(mqttConfig.Options.Keepalive)
		}

		optionsCache[server.Label] = serverOptions
	}

	client = &Client{
		clientCache:  make(map[string]mqtt.Client),
		optionsCache: optionsCache,
	}

	return
}
