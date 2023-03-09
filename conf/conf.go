package conf

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Aws struct {
		Region string `json:"region"`
	} `json:"aws"`
}

const APP_CONF_PREFIX = "ASE"

func LoadConfig() (Config, error) {
	var conf Config
	err := envconfig.Process(APP_CONF_PREFIX, &conf)

	return conf, err
}
