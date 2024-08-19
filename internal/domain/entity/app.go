package entity

type App struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Secret string `db:"secret"`
}
