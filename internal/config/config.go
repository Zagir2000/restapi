package config

import (
	"restapi/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"port"`
	Listen  struct {
		Type string `yaml:"type" env-default:"port"`
		Host string `yaml:"host" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
}

var DefaultConfigPath = "/root/restapi/config.yml"
var instance *Config
var once sync.Once

func GetConfig() *Config {
	/*Этот код выполним один раз, когда второй раз вызовут метод конфиг, то метод конфиг не будет выполняться,
	а будет возращаться instance(синглтоны не удобны для тестирования)*/
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(DefaultConfigPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			if err != nil {
				logger.Info(help)
				logger.Fatal(err)
			}

		}
	})
	return instance
}
