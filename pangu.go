package mqtt

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Musts(newMqtt)
}
