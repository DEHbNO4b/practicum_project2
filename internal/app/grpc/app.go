package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/DEHbNO4b/practicum_project2/internal/grpc/keeper"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	srv := grpc.NewServer()
	keeper.Register(srv)
	return &App{log: log, gRPCServer: srv, port: port}
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}
func (a *App) run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server")

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
