package mqtt

import (
	`github.com/pangum/pangu`
)

func init() {
	if err := pangu.New().Provides(newMqtt); nil != err {
		panic(err)
	}
}
