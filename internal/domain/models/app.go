package models

type App struct {
	id     int64
	name   string
	secret string
}

func (a *App) ID() int64 {
	return a.id
}

func (a *App) Name() string {
	return a.name
}

func (a *App) Secret() string {
	return a.secret
}
func (a *App) SetId(id int64) {
	a.id = id
}
func (a *App) SetName(name string) {
	a.name = name
}
func (a *App) SetSecret(secret string) {
	a.secret = secret
}
