package models

type User struct {
	id       int64
	login    string
	passHash string
}

func (u *User) ID() int64 {
	return u.id
}
func (u *User) Login() string {
	return u.login
}
func (u *User) PassHash() string {
	return u.passHash
}

func (u *User) SetID(id int64) {
	u.id = id
}
func (u *User) SetLogin(l string) {

	u.login = l
}
func (u *User) SetPassHash(ph string) {
	u.passHash = ph
}
