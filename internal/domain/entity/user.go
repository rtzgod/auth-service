package entity

type User struct {
	Id       int64  `db:"id"`
	Email    string `db:"email"`
	PassHash []byte `db:"pass_hash"`
}
