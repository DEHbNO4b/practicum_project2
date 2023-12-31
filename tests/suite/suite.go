package suite

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	myjwt "github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	certPath = "../certs/cert.pem"
)

type Suite struct {
	*testing.T
	Cfg        config.ClientConfig
	AuthClient pb.GophKeeperClient
	JWTClient  pb.GophKeeperClient
	ctx        context.Context
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadClientCfg()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.FileCfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	creds, err := credentials.NewClientTLSFromFile(certPath, "localhost")
	if err != nil {
		t.Fatalf("unable to read certFile %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.DialContext(
		ctx,
		cfg.FileCfg.GRPC.Host+":"+strconv.Itoa(cfg.FileCfg.GRPC.Port),
		opts...,
	)
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)

	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: pb.NewGophKeeperClient(conn),
		ctx:        ctx,
	}

}

func (s *Suite) MakeJWTClient(token string) error {

	cfg := config.MustLoadClientCfg()

	// jwtCreds, err := oauth.NewJWTAccessFromKey([]byte(token))
	// if err != nil {
	// 	return fmt.Errorf("failed to create JWT credentials: %v", err)
	// }
	jwtCreds := myjwt.JwtCredentials{
		Token: token, // Замените переменной актуальным токеном JWT
	}

	creds, err := credentials.NewClientTLSFromFile(certPath, "localhost")
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(jwtCreds),
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.DialContext(
		s.ctx,
		cfg.FileCfg.GRPC.Host+":"+strconv.Itoa(cfg.FileCfg.GRPC.Port),
		opts...,
	)
	if err != nil {
		return fmt.Errorf("failed to make dial conn: %v", err)

	}
	s.JWTClient = pb.NewGophKeeperClient(conn)

	return nil
}
