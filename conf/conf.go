package conf

import "github.com/kelseyhightower/envconfig"

type Config struct {
	WorkerInterval int `json:"workerInterval"`
	Aws            struct {
		Region string `json:"region"`
	} `json:"aws"`
}

const APP_CONF_PREFIX = "ASE"

func LoadConfig() (Config, error) {
	var conf Config
	err := envconfig.Process(APP_CONF_PREFIX, &conf)

	if conf.WorkerInterval == 0 {
		conf.WorkerInterval = 60
	}

	return conf, err
}
