package app

import (
	"github.com/jmoiron/sqlx"
	grpcserver "github.com/rtzgod/auth-service/internal/app/grpc"
	"github.com/rtzgod/auth-service/internal/repository/postgres"
	"github.com/rtzgod/auth-service/internal/services"
	"log/slog"
	"time"
)

type App struct {
	Log        *slog.Logger
	GRPCServer *grpcserver.Server
	Port       int
}

func NewApp(db *sqlx.DB, log *slog.Logger, port int, tokenTTL time.Duration) *App {

	repo := postgres.NewRepository(db)

	service := services.NewAuthService(log, repo, tokenTTL)

	gRPCServer := grpcserver.NewServer(log, port, service)

	return &App{
		Log:        log,
		GRPCServer: gRPCServer,
		Port:       port,
	}
}
