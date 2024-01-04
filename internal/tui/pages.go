package tui

func setPages(app *App) {
	app.Pages.AddPage("login form", app.AuthForm, true, true)
	app.Pages.AddPage("save data", app.SaveData, true, false)
	app.Pages.AddPage("show data", app.ShowData, true, false)
	app.Pages.AddPage("info", app.Info, true, false)
}
