package keeper

import (
	"context"
	"errors"
	"log/slog"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/services/auth"

	// pbauth "github.com/DEHbNO4b/practicum_project2/proto/gen/auth/proto"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, login string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, login string, password string) (userID int64, err error)
}

type Keeper interface {
	SaveLogPass(ctx context.Context, lp models.LogPassData) error
	LogPass(ctx context.Context, id int64) ([]models.LogPassData, error)

	SaveText(ctx context.Context, lp models.TextData) error
	TextData(ctx context.Context, id int64) ([]models.TextData, error)

	SaveBinary(ctx context.Context, lp models.BinaryData) error
	BinaryData(ctx context.Context, id int64) ([]models.BinaryData, error)

	SaveCard(ctx context.Context, lp models.Card) error
	CardData(ctx context.Context, id int64) ([]models.Card, error)
}
type ServerApi struct {
	log *slog.Logger
	pb.UnimplementedGophKeeperServer
	auth   Auth
	keeper Keeper
}

func Register(
	log *slog.Logger,
	srv *grpc.Server,
	auth Auth,
	keeper Keeper,
) {
	pb.RegisterGophKeeperServer(srv, &ServerApi{
		log:    log,
		auth:   auth,
		keeper: keeper,
	})
}

func (s *ServerApi) Register(ctx context.Context, req *pb.AuthInfo) (*pb.RegisterResponse, error) {

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
func (s *ServerApi) Login(ctx context.Context, req *pb.AuthInfo) (*pb.LoginResponse, error) {

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

// keeper handlers implementation
func (s *ServerApi) SaveLogPass(ctx context.Context, req *pb.SaveLogPassRequest) (*pb.SaveLogPassResponse, error) {

	op := "keeper/SaveLogPass"

	log := s.log.With(slog.String("op", op))

	log.Info("attemting to save log-pass data")

	res := pb.SaveLogPassResponse{}

	lpd := models.LogPassData{}
	lpd.SetLogin(req.GetLogin())
	lpd.SetPass(req.GetPassword())

	err := s.keeper.SaveLogPass(ctx, lpd)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
func validateLogin(req *pb.AuthInfo) error {

	if req.GetLogin() == "" {
		return status.Error(codes.InvalidArgument, "login is required")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}
func validateRegister(req *pb.AuthInfo) error {

	if req.GetLogin() == "" {
		return status.Error(codes.InvalidArgument, "login is required")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}
