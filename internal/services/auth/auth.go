package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/jwt"
	"github.com/DEHbNO4b/practicum_project2/internal/lib/logger/sl"
	"github.com/DEHbNO4b/practicum_project2/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type UserSaver interface {
	SaveUser(ctx context.Context, login string, passHash []byte) (uid int64, err error)
	Close()
}

type UserProvider interface {
	User(ctx context.Context, login string) (models.User, error)
	Close()
}

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	app         models.App
	tokenTTL    time.Duration
}

// New returns a new intance of Auth
func New(
	log *slog.Logger,
	us UserSaver,
	up UserProvider,
	app models.App,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    us,
		usrProvider: up,
		app:         app,
		tokenTTL:    tokenTTL,
	}
}
func (a *Auth) Stop() {
	a.usrProvider.Close()
	a.usrSaver.Close()
}
func (a *Auth) Login(ctx context.Context, login string, password string) (string, error) {
	const op = "auth.Login"
	log := a.log.With(slog.String("op", op))

	log.Info("attemtong to login user")

	user, err := a.usrProvider.User(ctx, login)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", slog.String("err", err.Error()))

			return "", fmt.Errorf("%s %w", op, ErrInvalidCredentials)
		}

		log.Error("unable to get user", sl.Err(err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash()), []byte(password)); err != nil {
		log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in succesfully")

	// app, err := a.appProvider.App(ctx)

	token, err := jwt.NewToken(user, a.app, a.tokenTTL)
	if err != nil {
		log.Error("unable to generate token", sl.Err(err))
		return "", err
	}

	return token, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, login string, password string) (int64, error) {

	const op = "auth.RegisterNewUser"
	log := a.log.With(slog.String("op", op))

	log.Info("register new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSaver.SaveUser(ctx, login, passHash)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("user already exists", sl.Err(err))

			return 0, fmt.Errorf("%s %w", op, ErrUserExists)
		}
		log.Error("unable to save user ", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
