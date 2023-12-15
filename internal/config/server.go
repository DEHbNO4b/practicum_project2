package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Env      string   `yaml:"env" env-default:"local"`
	DBconfig DBconfig `yaml:"dbconfig" env-required:"true"`
	GRPC     GRPCConfig
}

type DBconfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

func (db DBconfig) ToString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Database)
}

type GRPCConfig struct {
	Host    string        `yaml:"host"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoadServCfg() ServerConfig {
	path := filepath.FromSlash(fetchConfigPath())
	// f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	// if err != nil {
	// 	panic(err)

	// }
	// defer f.Close()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exists: " + path)
	}
	var serverCfg ServerConfig
	if err := cleanenv.ReadConfig(path, &serverCfg); err != nil {
		panic("cannot read config: " + err.Error())
	}
	return serverCfg
}
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "cfg", "./config/server.yaml", "path to server config yaml file")
	flag.Parse()

	return res
}
