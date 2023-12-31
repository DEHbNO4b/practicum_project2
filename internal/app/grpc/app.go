package grpcapp

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"

	// pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"github.com/DEHbNO4b/practicum_project2/internal/grpc/keeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingMetaData = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
	crtFile            = "./certs/cert.pem"
	keyFile            = "./certs/key.pem"
)

type App struct {
	log        *slog.Logger
	app        models.App
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	authService keeper.Auth,
	keeperService keeper.Keeper,
	port int,
	app models.App,

) (*App, error) {

	// srv := grpc.NewServer()
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Error("failed to load key pair %s", err)
		return nil, fmt.Errorf("unable to lead key pair from files %w", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	srv := grpc.NewServer(opts...)

	keeper.Register(
		srv,
		authService,
		keeperService,
	)

	return &App{log: log, gRPCServer: srv, port: port, app: app}, nil
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

	fmt.Println("in unary interceptor")
	fmt.Println(info.FullMethod)

	s := strings.Split(info.FullMethod, "/")

	if len(s) > 0 {
		switch s[len(s)-1] {
		case "Login":
			return handler(ctx, req)
		case "Register":
			return handler(ctx, req)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetaData
	}

	if !valid(md["authorization"]) {
		return nil, ErrInvalidToken
	}

	return handler(ctx, req)
}

func valid(auth []string) bool {
	if len(auth) < 1 {
		return false
	}
	token := strings.TrimPrefix(auth[0], "Bearer ")
	fmt.Println("got token in req: ", token)

	_, err := jwt.GetClaims(token)
	if err != nil {
		return false
	}

	return true
}
