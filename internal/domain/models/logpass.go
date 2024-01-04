package models

type LogPassData struct {
	userID int64
	login  string
	pass   string
	meta   string
}

func (lp *LogPassData) SetUserID(id int64) {
	lp.userID = id
}

func (lp *LogPassData) SetLogin(l string) {
	lp.login = l
}
func (lp *LogPassData) SetPass(p string) {
	lp.pass = p
}
func (lp *LogPassData) SetMeta(m string) {
	lp.meta = m
}

func (lp *LogPassData) UserID() int64 {
	return lp.userID
}
func (lp *LogPassData) Login() string {
	return lp.login
}
func (lp *LogPassData) Pass() string {
	return lp.pass
}
func (lp *LogPassData) Meta() string {
	return lp.meta
}

func (lp *LogPassData) String() string {
	return "login: " + lp.login + "\n" +
		"password: " + lp.pass + "\n" +
		"info: " + lp.meta
}
