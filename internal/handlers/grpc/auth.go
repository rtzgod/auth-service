package grpchandler

import (
	"context"
	"github.com/rtzgod/protos/gen/go/auth_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyValue = 0
)

func (h *GRPCHandler) SignIn(ctx context.Context, req *authv1.SignInRequest) (*authv1.SignInResponse, error) {

	if err := validateSignIn(req); err != nil {
		return nil, err
	}

	token, err := h.authService.SignIn(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &authv1.SignInResponse{
		Token: token,
	}, nil
}

func (h *GRPCHandler) SignUp(ctx context.Context, req *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {

	if err := validateSignUp(req); err != nil {
		return nil, err
	}

	userId, err := h.authService.SignUp(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.SignUpResponse{
		UserId: userId,
	}, nil
}

func (h *GRPCHandler) IsAdmin(ctx context.Context, req *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {

	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}

	isAdmin, err := h.authService.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func validateSignIn(req *authv1.SignInRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func validateSignUp(req *authv1.SignUpRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func validateIsAdmin(req *authv1.IsAdminRequest) error {
	if req.GetUserId() == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	return nil
}
