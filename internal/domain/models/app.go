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
