package main

import (
	"context"

	"github.com/DEHbNO4b/practicum_project2/internal/client"
	"github.com/DEHbNO4b/practicum_project2/internal/config"
)

func main() {

	//create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//create logger
	// log := setupLogger(cfg.Env)

	//read config
	cfg := config.MustLoadClientCfg()

	//create client
	client, err := client.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	switch cfg.Flags.LaunchMode {
	case "l":
		_, err := client.Login()
		if err != nil {
			return
		}

	case "r":
		_, err := client.Registert()
		if err != nil {
			return
		}
	}
}
