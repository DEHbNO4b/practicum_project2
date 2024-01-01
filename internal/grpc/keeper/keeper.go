package keeper

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	"github.com/DEHbNO4b/practicum_project2/internal/services/auth"

	// pbauth "github.com/DEHbNO4b/practicum_project2/proto/gen/auth/proto"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
// SaveLogPass is a handle to save log-pass data

func (s *ServerApi) SaveLogPass(ctx context.Context, req *pb.LogPassData) (*pb.Empty, error) {

	res := pb.Empty{}
	// op := "keeper/SaveLogPass"
	md, _ := metadata.FromIncomingContext(ctx)

	claims, _ := jwt.GetClaims(strings.TrimPrefix(md["authorization"][0], "Bearer "))

	lpd := models.LogPassData{}
	lpd.SetLogin(req.GetLogin())
	lpd.SetPass(req.GetPassword())
	lpd.SetMeta(req.GetInfo())
	lpd.SetUserID(claims.UserID)

	err := s.keeper.SaveLogPass(ctx, lpd)
	if err != nil {
		return &res, status.Error(codes.Internal, "internal error")
	}

	return &res, nil
}

// SaveLogPass is a handle to save log-pass data

func (s *ServerApi) SaveText(ctx context.Context, req *pb.TextData) (*pb.Empty, error) {

	res := pb.Empty{}

	md, _ := metadata.FromIncomingContext(ctx)

	claims, _ := jwt.GetClaims(strings.TrimPrefix(md["authorization"][0], "Bearer "))

	td := models.TextData{}
	td.SetText(req.GetText())
	td.SetMeta(req.GetInfo())
	td.SetUserID(claims.UserID)

	err := s.keeper.SaveText(ctx, td)
	if err != nil {
		return &res, status.Error(codes.Internal, "internal error")
	}

	return &res, nil
}

// SaveBinary is a handle to save log-pass data

func (s *ServerApi) SaveBinary(ctx context.Context, req *pb.BinaryData) (*pb.Empty, error) {

	res := pb.Empty{}

	md, _ := metadata.FromIncomingContext(ctx)

	claims, _ := jwt.GetClaims(strings.TrimPrefix(md["authorization"][0], "Bearer "))

	bd := models.BinaryData{}
	bd.SetData(req.GetData())
	bd.SetMeta(req.GetInfo())
	bd.SetUserID(claims.UserID)

	err := s.keeper.SaveBinary(ctx, bd)
	if err != nil {
		return &res, status.Error(codes.Internal, "internal error")
	}

	return &res, nil
}

func (s *ServerApi) ShowData(ctx context.Context, req *pb.Empty) (*pb.Data, error) {

	res := pb.Data{}

	md, _ := metadata.FromIncomingContext(ctx)

	claims, _ := jwt.GetClaims(strings.TrimPrefix(md["authorization"][0], "Bearer "))

	lpd, err := s.keeper.LogPass(ctx, claims.UserID)
	if err != nil {
		return &res, status.Error(codes.Internal, "internal error")
	}

	for _, el := range lpd {
		res.Lpd = append(res.Lpd, domainLogPassToProto(el))
	}

	td, err := s.keeper.TextData(ctx, claims.UserID)
	if err != nil {
		return &res, status.Error(codes.Internal, "internal error")
	}
	for _, el := range td {
		res.Td = append(res.Td, domainTextToProto(el))
	}
	bd, err := s.keeper.BinaryData(ctx, claims.UserID)
	if err != nil {
		return &res, status.Error(codes.Internal, "internal error")
	}
	for _, el := range bd {
		res.Bd = append(res.Bd, domainBinaryToProto(el))
	}

	return &res, nil
}
