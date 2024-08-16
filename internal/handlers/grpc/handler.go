package grpchandler

import (
	authv1 "github.com/rtzgod/protos/gen/go/auth_service"
	"google.golang.org/grpc"
)

type GRPCHandler struct {
	authv1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	authv1.RegisterAuthServer(gRPC, &GRPCHandler{})
}
