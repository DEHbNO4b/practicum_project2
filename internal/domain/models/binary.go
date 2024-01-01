package models

type BinaryData struct {
	userID int64
	data   []byte
	meta   string
}

func (b *BinaryData) SetUserID(id int64) {
	b.userID = id
}

func (b *BinaryData) SetData(data []byte) {
	b.data = make([]byte, 0, 20)
	b.data = append(b.data, data...)

}

func (b *BinaryData) SetMeta(meta string) {
	b.meta = meta
}

func (b *BinaryData) UserID() int64 {
	return b.userID
}

func (b *BinaryData) Data() []byte {
	ans := make([]byte, 0, len(b.data))
	ans = append(ans, b.data...)
	return ans
}

func (b *BinaryData) Meta() string {
	return b.meta
}
