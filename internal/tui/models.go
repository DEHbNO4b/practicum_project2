package tui

type userInfo struct {
	login    string
	password string
}

type LogPass struct {
	Login string
	Pass  string
	Info  string
}

type Card struct {
	CardID string
	Pass   string
	Date   string
	Info   string
}
type TextData struct {
	Text string
	Meta string
}

type BinaryData struct {
	Data string
	Info string
}
