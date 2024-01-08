package models

type TextData struct {
	userID int64
	text   string
	meta   string
}

func (t *TextData) SetUserID(id int64) {
	t.userID = id
}
func (t *TextData) SetText(text string) {
	t.text = text
}
func (t *TextData) SetMeta(meta string) {
	t.meta = meta
}

func (t *TextData) UserID() int64 {
	return t.userID
}
func (t *TextData) Text() string {
	return t.text
}
func (t *TextData) Meta() string {
	return t.meta
}

func (t *TextData) String() string {

	str := "text: " + t.text + "\n"
	str += "info: " + t.meta + "\n"

	return str
}
