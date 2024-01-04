package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	myjwt "github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var (
	hostname = "localhost"
	crtFile  = "./certs/client_cert.pem"
	keyFile  = "./certs/client_key.pem"
	caFile   = "./certs/ca_cert.pem"
)

var ErrNotAuthorized = errors.New("not authorized")

type GophClient struct {
	Ctx        context.Context
	AuthClient pb.GophKeeperClient
	JWTClient  pb.GophKeeperClient
	Cfg        config.ClientConfig
	Token      string
	Log        *slog.Logger
}

func New(ctx context.Context, cfg config.ClientConfig) (*GophClient, error) {

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to load client key pair %w", err)
	}

	certPool := x509.NewCertPool()

	ca, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("unable to load ca certificate %w", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, fmt.Errorf("unable to append ca certs %w", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostname,
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
		})),
	}
	conn, err := grpc.DialContext(
		ctx,
		cfg.GRPC.Host+":"+strconv.Itoa(cfg.GRPC.Port),
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("grpc server connection failed: %w", err)
	}

	client := pb.NewGophKeeperClient(conn)

	pbClient := GophClient{
		Ctx:        ctx,
		AuthClient: client,
		Cfg:        cfg,
	}
	return &pbClient, nil
}

func (g *GophClient) MakeJWTClient() error {

	if g.Token == "" {
		return errors.New("emty token")
	}

	cfg := config.MustLoadClientCfg()

	jwtCreds := myjwt.JwtCredentials{
		Token: g.Token,
	}

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return err
	}
	certPool := x509.NewCertPool()

	ca, err := os.ReadFile(caFile)
	if err != nil {
		return err
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(jwtCreds),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostname,
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
		})),
	}

	conn, err := grpc.DialContext(
		g.Ctx,
		cfg.GRPC.Host+":"+strconv.Itoa(cfg.GRPC.Port),
		opts...,
	)
	if err != nil {
		return fmt.Errorf("failed to make dial conn: %v", err)

	}
	g.JWTClient = pb.NewGophKeeperClient(conn)

	return nil
}

func (g *GophClient) Login(login, pass string) (models.User, error) {

	req := pb.AuthInfo{
		Login:    login,
		Password: pass,
	}

	u := models.User{}

	res, err := g.AuthClient.Login(g.Ctx, &req)

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.InvalidArgument {
				fmt.Println("wrong login or password, please try again")
				return u, err
			} else {
				// в остальных случаях выводим код ошибки в виде строки и сообщение
				fmt.Println(status.Code(), status.Message())
				return u, err
			}
		} else {
			fmt.Printf("Не получилось распарсить ошибку %v", err)
			return u, err
		}
	}

	u.SetLogin(res.GetName())

	g.Token = res.GetToken()

	err = g.MakeJWTClient()

	if err != nil {
		return u, err
	}

	return u, nil
}

func (g *GophClient) SignUp(login, pass string) error {

	_, err := g.AuthClient.Register(g.Ctx, &pb.AuthInfo{
		Login:    login,
		Password: pass,
	})

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.AlreadyExists {
				fmt.Println("user with this login already exists")
				return err
			} else {
				// в остальных случаях выводим код ошибки в виде строки и сообщение
				fmt.Println(status.Code(), status.Message())
				return err
			}
		} else {
			fmt.Printf("Не получилось распарсить ошибку %v", err)
			return err
		}
	}

	return nil
}

func (g *GophClient) SaveLogPass(ctx context.Context, lp *models.LogPassData) error {
	if g.JWTClient == nil {
		g.MakeJWTClient()
	}

	_, err := g.JWTClient.SaveLogPass(ctx, &pb.LogPassData{
		Login:    lp.Login(),
		Password: lp.Pass(),
		Info:     lp.Meta(),
	})
	return err
}

func (g *GophClient) SaveCard(ctx context.Context, c *models.Card) error {
	if g.JWTClient == nil {
		g.MakeJWTClient()
	}
	_, err := g.JWTClient.SaveCard(ctx, &pb.CardData{
		CardID: string(c.CardID()),
		Pass:   c.Pass(),
		Date:   c.Date(),
		Info:   c.Meta(),
	})
	return err
}

func (g *GophClient) SaveText(ctx context.Context, t *models.TextData) error {
	if g.JWTClient == nil {
		g.MakeJWTClient()
	}
	_, err := g.JWTClient.SaveText(ctx, &pb.TextData{
		Text: t.Text(),
		Info: t.Meta(),
	})
	return err
}

func (g *GophClient) SaveBinary(ctx context.Context, b *models.BinaryData) error {
	if g.JWTClient == nil {
		g.MakeJWTClient()
	}
	_, err := g.JWTClient.SaveBinary(ctx, &pb.BinaryData{
		Data: b.Data(),
		Info: b.Meta(),
	})
	return err
}
