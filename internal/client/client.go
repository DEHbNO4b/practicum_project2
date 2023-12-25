package client

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type GophClient struct {
	Ctx    context.Context
	Client pb.KeeperClient
	Cfg    config.ClientConfig
	Token  string
	Log    *slog.Logger
}

func New(ctx context.Context, cfg config.ClientConfig) (*GophClient, error) {

	conn, err := grpc.Dial(
		cfg.FileCfg.GRPC.Host+":"+strconv.Itoa(cfg.FileCfg.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	client := pb.NewKeeperClient(conn)

	pbClient := GophClient{
		Ctx:    ctx,
		Client: client,
		Cfg:    cfg,
	}
	return &pbClient, nil
}

func (g *GophClient) Login() (*pb.LoginResponse, error) {
	var (
		login string
		pass  string
	)
	fmt.Println("Enter login:")
	fmt.Scan(&login)
	fmt.Println("enter password")
	fmt.Scan(&pass)

	req := pb.LoginRequest{
		Login:    login,
		Password: pass,
	}

	res, err := g.Client.Login(g.Ctx, &req)

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.InvalidArgument {
				fmt.Println("wrong login or password, please try again")
				return res, err
			} else {
				// в остальных случаях выводим код ошибки в виде строки и сообщение
				fmt.Println(status.Code(), status.Message())
				return res, err
			}
		} else {
			fmt.Printf("Не получилось распарсить ошибку %v", err)
			return res, err
		}
	}

	g.Token = res.GetToken()

	return res, nil
}

func (g *GophClient) Registert() (*pb.RegisterResponse, error) {
	var (
		login string
		pass1 string
		pass2 string
	)
	fmt.Println("Enter login:")
	fmt.Scan(&login)
	fmt.Println("enter password")
	fmt.Scan(&pass1)
	fmt.Println("repeite password")
	fmt.Scan(&pass2)
	if pass1 != pass2 {
		fmt.Printf("passwords must be equal, please try again")
		return nil, errors.New("passwords not equal")
	}

	res, err := g.Client.Register(g.Ctx, &pb.RegisterRequest{
		Login:    login,
		Password: pass1,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.AlreadyExists {
				fmt.Println("user with this login already exists")
				return res, err
			} else {
				// в остальных случаях выводим код ошибки в виде строки и сообщение
				fmt.Println(status.Code(), status.Message())
				return res, err
			}
		} else {
			fmt.Printf("Не получилось распарсить ошибку %v", err)
			return res, err
		}
	}

	return res, nil
}
