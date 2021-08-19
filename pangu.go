package mqtt

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(newMqtt); nil != err {
		panic(err)
	}
}
