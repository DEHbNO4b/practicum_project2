package main

import (
	"context"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	"github.com/DEHbNO4b/practicum_project2/internal/grpc/client"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/logger/sl"
	"github.com/DEHbNO4b/practicum_project2/internal/tui"
)

func main() {

	//create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//read config
	cfg := config.MustLoadClientCfg()

	//create logger
	log := sl.SetupLogger(cfg.Env)

	//create client
	client, err := client.New(ctx, cfg, log)
	if err != nil {
		panic(err)
	}

	app := tui.New(ctx, client)

	if err := app.App.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
