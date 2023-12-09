package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/DEHbNO4b/practicum_project2/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, dsn string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{GRPCSrv: grpcApp}

}
