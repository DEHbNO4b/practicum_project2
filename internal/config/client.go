package config

import (
	"flag"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type ClientConfig struct {
	Env   string `yaml:"env" env-default:"local"`
	GRPC  GRPCConfig
	Flags ClientFlags
}

type ClientFlags struct {
	Path string
}

var (
	Flags  ClientFlags
	Config ClientConfig
	once   sync.Once
)

func MustLoadClientCfg() ClientConfig {

	// path := filepath.FromSlash(fetchConfigPath())
	once.Do(func() {
		parseFlags()

		if _, err := os.Stat(Flags.Path); os.IsNotExist(err) {
			panic("config file is not exists: " + Flags.Path)
		}

		if err := cleanenv.ReadConfig(Flags.Path, &Config); err != nil {
			panic("cannot read file configs: " + err.Error())
		}

		Config.Flags = Flags
	})

	return Config
}

func parseFlags() {

	flag.StringVar(&Flags.Path, "cfg", "./config/client.yaml", "path to config yaml file")

	flag.Parse()

}
