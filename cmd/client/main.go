package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	time.Sleep(3 * time.Second)
	cfg := config.MustLoadClientCfg()

	conn, err := grpc.Dial(
		cfg.GRPC.Host+":"+strconv.Itoa(cfg.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	client := pb.NewKeeperClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := pb.LoginRequest{
		Login:    "first",
		Password: "best_password",
	}
	response, err := client.Login(ctx, &req)
	if err != nil {
		fmt.Println("have got error from login request", err)
	}
	fmt.Printf("login response: %+v\n", response)
}
