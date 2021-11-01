package mqtt

import (
	`time`

	`github.com/eclipse/paho.mqtt.golang`
)

func setString(setter func(value string) *mqtt.ClientOptions, value string, null string) {
	if `` != value {
		setter(value)
	} else {
		setter(null)
	}
}

func setDuration(setter func(value time.Duration) *mqtt.ClientOptions, value time.Duration, null time.Duration) {
	if 0 != value {
		setter(value)
	} else {
		setter(null)
	}
}
