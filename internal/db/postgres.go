package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgres(url string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to postgres"))
	}
	err = db.Ping()
	if err != nil {
		panic(errors.Wrap(err, "failed to check connection with postgres"))
	}
	return db
}
