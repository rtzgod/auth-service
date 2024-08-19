package postgres

import (
	"context"
	"fmt"
	"github.com/rtzgod/auth-service/internal/domain/entity"
)

const (
	usersTable  = "users"
	appsTable   = "apps"
	adminsTable = "admins"
)

func (r *Repository) SaveUser(ctx context.Context, email string, passHash []byte) (id int64, err error) {
	const op = "repository.postgres.SaveUser"

	query := fmt.Sprintf("insert into %s(email, pass_hash) values ($1, $2) returning id", usersTable)

	err = r.db.Get(&id, query, email, passHash)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *Repository) User(ctx context.Context, email string) (user entity.User, err error) {
	const op = "repository.postgres.User"

	query := fmt.Sprintf("select * from %s where email=$1", usersTable)

	err = r.db.Get(&user, query, email)
	if err != nil {
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *Repository) App(ctx context.Context, appId int) (app entity.App, err error) {
	const op = "repository.postgres.App"

	query := fmt.Sprintf("select * from %s where id=$1", appsTable)

	err = r.db.Get(&app, query, appId)
	if err != nil {
		return entity.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}

func (r *Repository) IsAdmin(ctx context.Context, userId int64) (isAdmin bool, err error) {
	panic("implement me")
}
