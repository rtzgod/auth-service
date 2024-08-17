package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/rtzgod/auth-service/internal/pkg/jwt"
	"github.com/rtzgod/auth-service/internal/repository"
	"github.com/rtzgod/auth-service/pkg/sl"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidAppId       = errors.New("invalid app id")
	ErrUserExists         = errors.New("user already exists")
)

func (a *AuthService) SignIn(ctx context.Context, email, password string, appId int) (token string, err error) {
	const op = "AuthService.SignIn"

	log := a.log.With(slog.String("op", op))

	log.Info("Attempting to sign in user")

	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			log.Warn("User not found", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Error("failed to get user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := a.AppProvider.App(ctx, appId)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			log.Warn("App not found", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidAppId)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	token, err = jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (a *AuthService) SignUp(ctx context.Context, email, password string) (userId int64, err error) {
	const op = "AuthService.SignUp"

	log := a.log.With(slog.String("op", op))

	log.Info("Signing up user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.userSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			log.Warn("User already exists", sl.Err(err))

			return 0, fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		log.Error("failed to save user", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return id, nil
}

func (a *AuthService) IsAdmin(ctx context.Context, userId int64) (isAdmin bool, err error) {
	const op = "AuthService.IsAdmin"

	log := a.log.With(slog.String("op", op))

	log.Info("Checking if user is an admin")

	isAdmin, err = a.userProvider.IsAdmin(ctx, userId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			log.Warn("User not found", sl.Err(err))

			return false, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		log.Error("failed to check if user is an admin", sl.Err(err))

		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Checking if user is an admin", slog.Bool("isAdmin", isAdmin))

	return isAdmin, nil
}
