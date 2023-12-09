package main

import (
	"log/slog"
	"os"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
)

func main() {
	// TODO: get config
	cfg := config.MustLoad()

	// TODO: logger setup
	log := setupLogger(cfg.Env)

	log.Info("starting app")
	log.Info("cfg", slog.Any("", cfg))
	//TODO: creat storage

	//TODO: crate server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
