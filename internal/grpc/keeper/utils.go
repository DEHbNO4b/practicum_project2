package keeper

import (
	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/keeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func domainLogPassToProto(lp models.LogPassData) *pb.LogPassData {
	return &pb.LogPassData{
		Login:    lp.Login(),
		Password: lp.Pass(),
		Info:     lp.Meta(),
	}
}
func domainTextToProto(td models.TextData) *pb.TextData {
	return &pb.TextData{
		Text: td.Text(),
		Info: td.Meta(),
	}
}
