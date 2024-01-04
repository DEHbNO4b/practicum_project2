package tui

import (
	"fmt"
	"time"
)

func setAuthForm(app *App) {
	// app.AuthForm.SetBorder(true).SetTitle("Login")
	app.AuthForm.AddInputField("Имя пользователя", "", 20, nil, func(login string) {
		app.ClientInfo.login = login
	}).
		AddPasswordField("Пароль", "", 20, '*', func(pass string) {
			app.ClientInfo.password = pass
		}).
		AddButton("Логин", func() { // Обработка логина
			u, err := app.client.Login(app.ClientInfo.login, app.ClientInfo.password)
			if err != nil {
				fmt.Println("Game Over")
				app.App.Stop()

			}
			app.ClientInfo.login = u.Login()

			app.Pages.SwitchToPage("save data")

		}).
		AddButton("Регистрация", func() { // Обработка регистрации
			err := app.client.SignUp(app.ClientInfo.login, app.ClientInfo.password)
			if err != nil {
				fmt.Println("Game Over")
				time.Sleep(10 * time.Second)
				app.App.Stop()
			}

			u, err := app.client.Login(app.ClientInfo.login, app.ClientInfo.password)
			if err != nil {
				fmt.Println("Game Over")
				time.Sleep(10 * time.Second)
				app.App.Stop()
			}

			app.ClientInfo.login = u.Login()

			app.Pages.SwitchToPage("save data")
		}).
		AddButton("Выход", func() {
			app.App.Stop()
		})
}
