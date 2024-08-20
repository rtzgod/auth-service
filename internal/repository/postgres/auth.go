package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/rtzgod/auth-service/internal/domain/entity"
	"github.com/rtzgod/auth-service/internal/repository"
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
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, repository.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *Repository) User(ctx context.Context, email string) (user entity.User, err error) {
	const op = "repository.postgres.User"

	query := fmt.Sprintf("select * from %s where email=$1", usersTable)

	err = r.db.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
		}
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *Repository) App(ctx context.Context, appId int) (app entity.App, err error) {
	const op = "repository.postgres.App"

	query := fmt.Sprintf("select * from %s where id=$1", appsTable)

	err = r.db.Get(&app, query, appId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.App{}, fmt.Errorf("%s: %w", op, repository.ErrAppNotFound)
		}
		return entity.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}

func (r *Repository) IsAdmin(ctx context.Context, userId int64) (isAdmin bool, err error) {
	const op = "repository.postgres.IsAdmin"

	query := fmt.Sprintf("select * from %s where id=$1", adminsTable)

	err = r.db.Get(&isAdmin, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return isAdmin, nil
}
