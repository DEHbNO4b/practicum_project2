package keeper

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/logger/sl"
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
	res.Name = req.GetLogin()

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

// KEEPER INTERFACE IMPLEMENTATION
//
// SaveLogPass is a handle to save log-pass data
func (s *ServerApi) SaveLogPass(ctx context.Context, req *pb.LogPassData) (*pb.Empty, error) {

	res := pb.Empty{}
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

// SaveText is a handle to save text data

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

// SaveBinary is a handle to save binary data
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

// SaveCard is a handle to save card data
func (s *ServerApi) SaveCard(ctx context.Context, req *pb.CardData) (*pb.Empty, error) {

	res := pb.Empty{}

	md, _ := metadata.FromIncomingContext(ctx)

	claims, _ := jwt.GetClaims(strings.TrimPrefix(md["authorization"][0], "Bearer "))

	cd := models.Card{}
	cd.SetCardID([]rune(req.GetCardID()))
	cd.SetPass(req.GetPass())
	cd.SetDate(req.GetDate())
	cd.SetMeta(req.GetInfo())
	cd.SetUserID(claims.UserID)

	err := s.keeper.SaveCard(ctx, cd)
	if err != nil {
		s.log.Error("unable to save card data to db", err)
		return &res, status.Error(codes.Internal, "internal error")
	}

	return &res, nil
}

// ShowData is a handle to show all data types
func (s *ServerApi) ShowData(ctx context.Context, req *pb.Empty) (*pb.Data, error) {

	op := "grpc/keeper/ShowData"

	log := s.log.With(slog.String("op", op))

	res := pb.Data{}

	md, _ := metadata.FromIncomingContext(ctx)

	claims, _ := jwt.GetClaims(strings.TrimPrefix(md["authorization"][0], "Bearer "))

	lpd, err := s.keeper.LogPass(ctx, claims.UserID) // get log-pass data from keeper
	if err != nil {
		log.Error("unable to get LogPass data", sl.Err(err))
		return &res, status.Error(codes.Internal, "internal error")
	}

	for _, el := range lpd {
		res.Lpd = append(res.Lpd, domainLogPassToProto(el))
	}

	td, err := s.keeper.TextData(ctx, claims.UserID) // Get text data from keeper
	if err != nil {
		log.Error("unable to get Text data", sl.Err(err))
		return &res, status.Error(codes.Internal, "internal error")
	}
	for _, el := range td {
		res.Td = append(res.Td, domainTextToProto(el))
	}

	bd, err := s.keeper.BinaryData(ctx, claims.UserID) // Get binary data from keeper
	if err != nil {
		log.Error("unable to get binary data", sl.Err(err))
		return &res, status.Error(codes.Internal, "internal error")
	}
	for _, el := range bd {
		res.Bd = append(res.Bd, domainBinaryToProto(el))
	}

	cd, err := s.keeper.CardData(ctx, claims.UserID) // Get card data from keeper
	if err != nil {
		log.Error("unable to get card data", sl.Err(err))
		return &res, status.Error(codes.Internal, "internal error")
	}
	for _, el := range cd {
		res.Cd = append(res.Cd, domainCardToProto(el))
	}

	return &res, nil
}
