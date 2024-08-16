package grpchandler

import (
	"context"
	authv1 "github.com/rtzgod/protos/gen/go/auth_service"
	"google.golang.org/grpc"
)

type AuthService interface {
	SignIn(ctx context.Context, email, password string, appId int) (token string, err error)
	SignUp(ctx context.Context, email, password string) (userId int64, err error)
	IsAdmin(ctx context.Context, userId int) (isAdmin bool, err error)
}

type GRPCHandler struct {
	authv1.UnimplementedAuthServer
	authService AuthService
}

func AddHandler(gRPC *grpc.Server, authService AuthService) {
	authv1.RegisterAuthServer(gRPC, &GRPCHandler{authService: authService})
}
