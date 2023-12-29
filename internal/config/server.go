package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var servOnce sync.Once
var ServerCfg ServerConfig

type ServerConfig struct {
	Env       string        `yaml:"env" env-default:"local"`
	TokenTTL  time.Duration `yaml:"tokenTTL" env-default:"1h"`
	SecretKey string        `yaml:"secret" env-required:"true"`
	DBconfig  DBconfig      `yaml:"dbconfig" env-required:"true"`
	GRPC      GRPCConfig
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

	servOnce.Do(func() {
		path := filepath.FromSlash(fetchConfigPath())

		MustLoadByPath(path)
	})

	return ServerCfg

}
func fetchConfigPath() string {

	var res string

	flag.StringVar(&res, "cfg", "./config/server.yaml", "path to server config yaml file")
	flag.Parse()

	return res
}
func MustLoadByPath(path string) ServerConfig {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exists: " + path)
	}

	if err := cleanenv.ReadConfig(path, &ServerCfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return ServerCfg
}
