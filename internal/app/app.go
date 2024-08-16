package app

import (
	grpcserver "github.com/rtzgod/auth-service/internal/app/grpc"
	"log/slog"
)

type App struct {
	Log        *slog.Logger
	GRPCServer *grpcserver.Server
	Port       int
}

func NewApp(log *slog.Logger, port int) *App {
	gRPCServer := grpcserver.NewServer(log, port)

	return &App{
		Log:        log,
		GRPCServer: gRPCServer,
		Port:       port,
	}
}
