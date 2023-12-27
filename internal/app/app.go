package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/DEHbNO4b/practicum_project2/internal/app/grpc"
	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/services/auth"
	"github.com/DEHbNO4b/practicum_project2/internal/services/keeper"
	"github.com/DEHbNO4b/practicum_project2/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, dsn string, tokenTTL time.Duration) *App {

	storage, err := postgres.New(dsn)
	if err != nil {
		panic(err)
	}

	app := models.App{}
	app.SetId(1)
	app.SetName("auth")
	app.SetSecret("secret_string")

	authService := auth.New(
		log,
		storage,
		storage,
		app,
		tokenTTL,
	)

	keeperService := keeper.New(
		log,
		storage,
		storage,
		storage,
		storage,
	)

	grpcApp := *grpcapp.New(log, authService, keeperService, grpcPort)

	return &App{
		GRPCSrv: &grpcApp,
	}
}
