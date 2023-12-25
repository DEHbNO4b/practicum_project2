package postgres

type User struct {
	Id       int64
	Login    string
	PassHash string
}

type LogPassData struct {
	UserID int64
	Login  string
	Pass   string
	Meta   string
}
