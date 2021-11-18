package mqtt

import (
	`os`

	`github.com/google/uuid`
	`github.com/pangum/pangu`
	`github.com/rs/xid`
)

type panguConfig struct {
	// MQTT的配置信息
	Mqtt config `json:"mqtt" yaml:"mqtt" xml:"mqtt" toml:"mqtt" validate:"required"`
}

func (pc *panguConfig) load(config *pangu.Config) (err error) {
	envs := map[string]string{
		`UUID`: uuid.NewString(),
		`XID`:  xid.New().String(),
	}
	if envs[`HOSTNAME`], err = os.Hostname(); nil != err {
		return
	}

	// 注入环境变量
	for env, value := range envs {
		if err = os.Setenv(env, value); nil != err {
			return
		}
	}
	// 加载配置文件
	err = config.Load(pc)

	return
}
