package postgres

import (
	"context"
	"github.com/rtzgod/auth-service/internal/domain/entity"
)

func (r *Repository) SaveUser(ctx context.Context, email string, passHash []byte) (id int64, err error) {
	return 1, nil
}

func (r *Repository) User(ctx context.Context, email string) (user entity.User, err error) {
	panic("implement me")
}

func (r *Repository) IsAdmin(ctx context.Context, userId int64) (isAdmin bool, err error) {
	panic("implement me")
}

func (r *Repository) App(ctx context.Context, appId int) (app entity.App, err error) {
	panic("implement me")
}
