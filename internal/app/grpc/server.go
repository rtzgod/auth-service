package grpcserver

import (
	"fmt"
	grpchandler "github.com/rtzgod/auth-service/internal/handlers/grpc"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewServer(log *slog.Logger, port int, authService grpchandler.AuthService) *Server {
	gRPCServer := grpc.NewServer()

	grpchandler.AddHandler(gRPCServer, authService)

	return &Server{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "grpcserver.Run"

	s.log.With(slog.String("op", op), slog.Int("port", s.port))

	s.log.Info("Starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("gRPC server in running", slog.String("addr", l.Addr().String()))

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Stop() {
	const op = "grpcserver.Stop"

	s.log.With(slog.String("op", op))
	s.log.Info("Stopping gRPC server", slog.Int("port", s.port))

	s.gRPCServer.GracefulStop()
}
