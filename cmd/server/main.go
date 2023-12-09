package main

import (
	"log/slog"
	"os"

	"github.com/DEHbNO4b/practicum_project2/internal/app"
	"github.com/DEHbNO4b/practicum_project2/internal/config"
)

func main() {
	// getting config
	cfg := config.MustLoad()

	//  logger setup
	log := setupLogger(cfg.Env)

	log.Info("starting app")
	log.Info("cfg", slog.Any("", cfg))

	application := app.New(log, cfg.GRPC.Port, cfg.DBconfig.ToString(), cfg.GRPC.Timeout)
	application.GRPCSrv.MustRun()

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
