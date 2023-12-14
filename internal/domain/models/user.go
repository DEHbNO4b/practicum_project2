package models

type User struct {
	login    string
	passHash string
}

func (u *User) Login() string {
	return u.login
}
func (u *User) PassHash() string {
	return u.passHash
}
