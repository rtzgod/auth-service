package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgres(host, port, user, password, dbname, sslmode string) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to postgres"))
	}
	err = db.Ping()
	if err != nil {
		panic(errors.Wrap(err, "failed to check connection with postgres"))
	}
	return db
}
