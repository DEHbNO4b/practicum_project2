package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/DEHbNO4b/practicum_project2/internal/grpc/keeper"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	authService keeper.Auth,
	keeperService keeper.Keeper,
	port int,

) *App {

	srv := grpc.NewServer()

	keeper.Register(
		srv,
		authService,
		keeperService,
	)

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

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	// if token != SecretToken {
	// 	return nil, status.Error(codes.Unauthenticated, "invalid token")
	// }

	claims, err := jwt.GetClaims(token)
	return handler(ctx, req)
}
