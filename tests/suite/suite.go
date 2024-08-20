package suite

import (
	"context"
	"github.com/rtzgod/auth-service/internal/config"
	authv1 "github.com/rtzgod/protos/gen/go/auth_service"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient authv1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()
}
