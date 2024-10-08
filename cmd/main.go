package main

import (
	"github.com/rtzgod/auth-service/internal/app"
	"github.com/rtzgod/auth-service/internal/config"
	"github.com/rtzgod/auth-service/internal/db"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting app", slog.String("env", cfg.Env))

	database := db.NewPostgres(cfg.Postgres.Url)

	// Applying migrations to db
	db.MigrateUp(database.DB)

	application := app.NewApp(database, log, cfg.GRPC.Port, cfg.GRPC.Timeout)

	go application.GRPCServer.MustRun()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sign := <-sigChan

	log.Info("stopping app,", slog.String("signal:", sign.String()))

	application.GRPCServer.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
