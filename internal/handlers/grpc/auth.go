package grpchandler

import (
	"context"
	authv1 "github.com/rtzgod/protos/gen/go/auth_service"
)

func (h *GRPCHandler) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	return &authv1.RegisterResponse{
		UserId: 1,
	}, nil
}

func (h *GRPCHandler) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	return &authv1.LoginResponse{
		Token: "valera",
	}, nil
}

func (h *GRPCHandler) IsAdmin(ctx context.Context, req *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	return &authv1.IsAdminResponse{
		IsAdmin: true,
	}, nil
}
