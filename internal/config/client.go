package config

import (
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type ClientConfig struct {
	Env  string `yaml:"env" env-default:"local"`
	GRPC GRPCConfig
}

func MustLoadClientCfg() ClientConfig {
	path := filepath.FromSlash(fetchConfigPath())

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exists: " + path)
	}

	var cfg ClientConfig
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return cfg
}
