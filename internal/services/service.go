package services

import (
	"context"
	"github.com/rtzgod/auth-service/internal/domain/entity"
	"log/slog"
	"time"
)

type AuthService struct {
	log          *slog.Logger
	userSaver    UserSaver
	userProvider UserProvider
	AppProvider  AppProvider
	tokenTTL     time.Duration
}

type Repository interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (id int64, err error)
	User(ctx context.Context, email string) (user entity.User, err error)
	IsAdmin(ctx context.Context, userId int64) (isAdmin bool, err error)
	App(ctx context.Context, appId int) (app entity.App, err error)
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (id int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (user entity.User, err error)
	IsAdmin(ctx context.Context, userId int64) (isAdmin bool, err error)
}

type AppProvider interface {
	App(ctx context.Context, appId int) (app entity.App, err error)
}

func NewAuthService(log *slog.Logger, repo Repository, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		log:          log,
		userSaver:    repo,
		userProvider: repo,
		AppProvider:  repo,
		tokenTTL:     tokenTTL,
	}
}
