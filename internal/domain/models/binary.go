package models

type BinaryData struct {
	userID int64
	data   []byte
	meta   string
}

func (b BinaryData) SetUserID(id int64) {
	b.userID = id
}

func (b BinaryData) SetData(data []byte) {
	b.data = data
}

func (b BinaryData) SetMeta(meta string) {
	b.meta = meta
}

func (b BinaryData) UserID() int64 {
	return b.userID
}

func (b BinaryData) Data() []byte {
	return b.data
}

func (b BinaryData) Meta() string {
	return b.meta
}
