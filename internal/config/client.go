package config

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type ClientConfig struct {
	FileCfg FileCfg
	Flags   ClientFlags
}

type FileCfg struct {
	Env  string `yaml:"env" env-default:"local"`
	GRPC GRPCConfig
}
type ClientFlags struct {
	LaunchMode string
	Path       string
}

var (
	Flags      ClientFlags
	FileConfig FileCfg
	Config     ClientConfig
	once       sync.Once
)

func MustLoadClientCfg() ClientConfig {

	// path := filepath.FromSlash(fetchConfigPath())
	once.Do(func() {
		fmt.Println("once do ")
		parseFlags()

		if _, err := os.Stat(Flags.Path); os.IsNotExist(err) {
			panic("config file is not exists: " + Flags.Path)
		}

		if err := cleanenv.ReadConfig(Flags.Path, &FileConfig); err != nil {
			panic("cannot read file configs: " + err.Error())
		}
		Config.FileCfg = FileConfig
		Config.Flags = Flags
	})

	return Config
}

func parseFlags() {

	flag.StringVar(&Flags.Path, "cfg", "./config/server.yaml", "path to config yaml file")
	flag.StringVar(&Flags.LaunchMode, "m", "./config/server.yaml", "path to config yaml file")

	flag.Parse()

}
