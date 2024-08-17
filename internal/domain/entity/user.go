package entity

type User struct {
	Id       int64
	Email    string
	PassHash []byte
}
