package keeper

import (
	"context"
	"errors"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/services/auth"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, login string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, login string, password string) (userID int64, err error)
}

type LogPassKeeper interface {
	SaveLogPass(ctx context.Context, lp models.LogPassData) error
	LogPass(ctx context.Context, id int64) ([]models.LogPassData, error)
}
type TextKeeper interface {
	SaveText(ctx context.Context, lp models.TextData) error
	TextData(ctx context.Context, id int64) ([]models.TextData, error)
}
type BinaryKeeper interface {
	SaveBinary(ctx context.Context, lp models.BinaryData) error
	BinaryData(ctx context.Context, id int64) ([]models.BinaryData, error)
}
type CardKeeper interface {
	SaveCard(ctx context.Context, lp models.Card) error
	CardData(ctx context.Context, id int64) ([]models.Card, error)
}
type ServerApi struct {
	pb.UnimplementedKeeperServer
	auth         Auth
	lpKeeper     LogPassKeeper
	textKeeper   TextKeeper
	binaryKeeper BinaryKeeper
	cardKeeper   CardKeeper
}

func Register(
	srv *grpc.Server,
	auth Auth,
	lpKeeper LogPassKeeper,
	textKeeper TextKeeper,
	binaryKeeper BinaryKeeper,
	cardKeeper CardKeeper,
) {
	pb.RegisterKeeperServer(srv, &ServerApi{
		auth:         auth,
		lpKeeper:     lpKeeper,
		textKeeper:   textKeeper,
		binaryKeeper: binaryKeeper,
		cardKeeper:   cardKeeper,
	})
}

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	res := pb.RegisterResponse{}

	if err := validateRegister(req); err != nil {
		return nil, err
	}

	id, err := s.auth.RegisterNewUser(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		if errors.Is(err, auth.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	res.UserId = id

	return &res, nil
}
func (s *ServerApi) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	res := pb.LoginResponse{}

	err := validateLogin(req)
	if err != nil {
		return nil, err
	}

	t, err := s.auth.Login(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {

		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid credentials")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	res.Token = t

	return &res, nil
}
func validateLogin(req *pb.LoginRequest) error {
	if req.GetLogin() == "" {
		return status.Error(codes.InvalidArgument, "login is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}
func validateRegister(req *pb.RegisterRequest) error {
	if req.GetLogin() == "" {
		return status.Error(codes.InvalidArgument, "login is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}
