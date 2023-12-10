package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/services/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserSaver interface {
	SaveUser(ctx context.Context, login string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, login string) (models.User, error)
}
type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
}

// New returns a new intance of Auth
func New(log *slog.Logger, us UserSaver, up UserProvider) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    us,
		usrProvider: up,
	}
}
func (a *Auth) Login(ctx context.Context, login string, password string) (string, error) {
	const op = "auth.Login"
	log := a.log.With(slog.String("op", op))

	log.Info("attemtong to login user")

	u,err:=a.usrProvider.User(ctx,login)
	if err!=nil{
		if errors.Is(err,storage.ErrUserNotFound)
		log.Warn("user not found",slog.StringValue("err",err.Error()))
	}
	
	return "", nil
}
func (a *Auth) RegisterNewUser(ctx context.Context, login string, password string) (int64, error) {
	const op = "auth.RegisterNewUser"
	log := a.log.With(slog.String("op", op))

	log.Info("register new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", slog.String("err", err.Error()))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSaver.SaveUser(ctx, login, passHash)
	if err != nil {
		log.Error("unable to save user ", slog.String("err", err.Error()))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
