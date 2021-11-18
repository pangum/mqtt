package mqtt

import (
	`crypto/tls`
	`net/url`

	`github.com/eclipse/paho.mqtt.golang`
	`github.com/pangum/logging`
	`github.com/pangum/pangu`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

func newMqtt(config *pangu.Config, logger *logging.Logger) (client *Client, err error) {
	getClient := func() *Client {
		return client
	}

	_config := new(panguConfig)
	if err = _config.load(config); nil != err {
		return
	}

	mqttConfig := _config.Mqtt
	// 加载默认连接
	brokerCache := make(map[string]broker)
	optionsCache := make(map[string]*mqtt.ClientOptions)
	serializerCache := make(map[string]serializer)
	if mqttConfig.Broker.validate() {
		_defaultOptions := mqtt.NewClientOptions()
		_defaultOptions.AddBroker(mqttConfig.Broker.best())
		_defaultOptions.SetUsername(mqttConfig.Options.Username)
		_defaultOptions.SetPassword(mqttConfig.Options.Password)
		_defaultOptions.SetKeepAlive(mqttConfig.Options.Keepalive)
		_defaultOptions.SetClientID(mqttConfig.Options.Clientid)
		if mqttConfig.Options.Will.Enabled {
			_defaultOptions.SetWill(
				mqttConfig.Options.Will.Topic,
				mqttConfig.Options.Will.Payload,
				byte(mqttConfig.Options.Will.Qos),
				mqttConfig.Options.Will.Retained,
			)
		}
		// 自动重连
		_defaultOptions.SetAutoReconnect(mqttConfig.Options.Reconnect.Auto)
		_defaultOptions.SetMaxReconnectInterval(mqttConfig.Options.Reconnect.Interval)
		_defaultOptions.SetResumeSubs(mqttConfig.Options.Reconnect.Resume)
		// 会话
		_defaultOptions.SetCleanSession(mqttConfig.Options.Session.Clean)
		// 重试
		_defaultOptions.SetConnectRetry(mqttConfig.Options.Retry.Enable)
		_defaultOptions.SetConnectRetryInterval(mqttConfig.Options.Retry.Interval)
		// 超时
		_defaultOptions.SetPingTimeout(mqttConfig.Options.Timeout.Ping)
		_defaultOptions.SetConnectTimeout(mqttConfig.Options.Timeout.Connect)
		_defaultOptions.SetWriteTimeout(mqttConfig.Options.Timeout.Write)
		// 处理器
		_defaultOptions.OnReconnecting = onReconnection(getClient, logger)
		_defaultOptions.OnConnectionLost = onConnectionLost(logger)
		_defaultOptions.OnConnect = onConnect(logger)
		_defaultOptions.OnConnectAttempt = onConnectAttempt(logger)

		optionsCache[defaultLabel] = _defaultOptions
		brokerCache[defaultLabel] = mqttConfig.Broker
		serializerCache[defaultLabel] = mqttConfig.Options.Serializer
	}

	// 加载带标签的服务器
	for _, _server := range mqttConfig.Brokers {
		if _server.Broker.validate() {
			continue
		}

		serverOptions := mqtt.NewClientOptions()
		serverOptions.AddBroker(_server.Broker.best())
		setString(serverOptions.SetUsername, _server.Options.Username, mqttConfig.Options.Username)
		setString(serverOptions.SetPassword, _server.Options.Password, mqttConfig.Options.Password)
		setDuration(serverOptions.SetKeepAlive, _server.Options.Keepalive, mqttConfig.Options.Keepalive)
		setString(serverOptions.SetClientID, _server.Options.Clientid, mqttConfig.Options.Clientid)
		if _server.Options.Will.Enabled {
			serverOptions.SetWill(
				_server.Options.Will.Topic,
				_server.Options.Will.Payload,
				byte(_server.Options.Will.Qos),
				_server.Options.Will.Retained,
			)
		}
		// 自动重连
		serverOptions.SetAutoReconnect(_server.Options.Reconnect.Auto)
		setDuration(serverOptions.SetMaxReconnectInterval, _server.Options.Reconnect.Interval, mqttConfig.Options.Reconnect.Interval)
		serverOptions.SetResumeSubs(_server.Options.Reconnect.Resume)
		// 会话
		serverOptions.SetCleanSession(_server.Options.Session.Clean)
		// 重试
		serverOptions.SetConnectRetry(_server.Options.Retry.Enable)
		serverOptions.SetConnectRetryInterval(_server.Options.Retry.Interval)
		// 超时
		setDuration(serverOptions.SetPingTimeout, _server.Options.Timeout.Ping, mqttConfig.Options.Timeout.Ping)
		setDuration(serverOptions.SetConnectTimeout, _server.Options.Timeout.Connect, mqttConfig.Options.Timeout.Connect)
		setDuration(serverOptions.SetWriteTimeout, _server.Options.Timeout.Write, mqttConfig.Options.Timeout.Write)
		// 处理器
		serverOptions.OnReconnecting = onReconnection(getClient, logger)
		serverOptions.OnConnectionLost = onConnectionLost(logger)
		serverOptions.OnConnect = onConnect(logger)
		serverOptions.OnConnectAttempt = onConnectAttempt(logger)

		optionsCache[_server.Label] = serverOptions
		brokerCache[_server.Label] = _server.Broker
		serializerCache[_server.Label] = _server.Options.Serializer
	}
	client = newClient(optionsCache, brokerCache, serializerCache, logger)

	return
}

func onConnectAttempt(logger *logging.Logger) func(broker *url.URL, tlsCfg *tls.Config) *tls.Config {
	return func(broker *url.URL, tlsCfg *tls.Config) *tls.Config {
		logger.Info(`尝试连接MQTT服务器`, field.Strings(`server`, broker.String()))

		return tlsCfg
	}
}

func onConnect(logger *logging.Logger) func(mqtt.Client) {
	return func(client mqtt.Client) {
		_options := client.OptionsReader()
		var msg string
		if client.IsConnected() {
			msg = `连接MQTT服务器成功`
		} else {
			msg = `连接MQTT服务器中`
		}

		logger.Info(
			msg,
			field.Strings(`urls`, servers(_options.Servers())...),
			field.String(`username`, _options.Username()),
			field.String(`clientid`, _options.ClientID()),
		)
	}
}

func onConnectionLost(logger *logging.Logger) func(mqtt.Client, error) {
	return func(client mqtt.Client, err error) {
		_options := client.OptionsReader()
		logger.Warn(
			`MQTT掉线`,
			field.Strings(`urls`, servers(_options.Servers())...),
			field.String(`username`, _options.Username()),
			field.String(`clientid`, _options.ClientID()),
			field.Error(err))
	}
}

func onReconnection(fun func() *Client, logger *logging.Logger) func(mqtt.Client, *mqtt.ClientOptions) {
	return func(mqtt mqtt.Client, options *mqtt.ClientOptions) {
		logger.Warn(
			`MQTT自动重连中`,
			field.Strings(`urls`, servers(options.Servers)...),
			field.String(`username`, options.Username),
			field.String(`clientid`, options.ClientID),
			field.Bool(`resume`, options.ResumeSubs),
		)

		if !options.ResumeSubs {
			return
		}

		client := fun()
		if nil == client {
			return
		}

		successes, fails, err := client.Resubscribe()
		fields := []gox.Field{
			field.Strings(`successes`, successes...),
			field.Strings(`fails`, fails...),
		}
		if nil != err {
			fields = append(fields, field.Error(err))
		}
		logger.Info(`恢复订阅关系`, fields...)
	}
}

func servers(urls []*url.URL) (servers []string) {
	for _, _url := range urls {
		servers = append(servers, _url.String())
	}

	return
}
