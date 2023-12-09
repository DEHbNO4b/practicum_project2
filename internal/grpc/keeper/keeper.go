package keeper

import (
	"context"

	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/proto"
	"google.golang.org/grpc"
)

type ServerApi struct {
	pb.UnimplementedKeeperServer
}

func Register(srv *grpc.Server) {
	pb.RegisterKeeperServer(srv, &ServerApi{})
}
func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res := pb.RegisterResponse{}
	return &res, nil
}
func (s *ServerApi) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res := pb.LoginResponse{}
	res.Token = "token string: best way to check grpc handler"
	return &res, nil
}
