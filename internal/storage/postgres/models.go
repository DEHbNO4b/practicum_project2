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

type TextData struct {
	UserID int64
	Text   string
	Meta   string
}

type BinaryData struct {
	UserID int64
	Data   []byte
	Meta   string
}

type Card struct {
	UserID int64
	CardID string
	Pass   string
	Date   string
	Meta   string
}
