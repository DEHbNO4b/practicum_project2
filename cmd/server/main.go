package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DEHbNO4b/practicum_project2/internal/app"
	"github.com/DEHbNO4b/practicum_project2/internal/config"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/logger/sl"
)

func main() {
	// getting config
	cfg := config.MustLoadServCfg()

	//  logger setup
	log := sl.SetupLogger(cfg.Env)

	log.Info("starting app")
	log.Info("cfg", slog.Any("", cfg))

	application, err := app.New(log, cfg.GRPC.Port, cfg.DBconfig.ToString(), cfg.GRPC.Timeout)
	if err != nil {
		panic(err)
	}
	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))
	application.GRPCSrv.Stop()

}
