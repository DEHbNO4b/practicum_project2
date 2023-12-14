package postgres

type User struct {
	Id       int64
	Login    string
	PassHash string
}
