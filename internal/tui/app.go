package tui

import (
	"fmt"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	clientInfo userInfo
)

type GophClient interface {
	SignUp(login, pass string) error
	Login(login, pass string) (models.User, error)
}

func App(client GophClient) *tview.Application {

	pages := tview.NewPages()

	greetingText := tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("Hello, " + clientInfo.login)

	app := tview.NewApplication()

	// Создаем форму для ввода данных
	AuthForm := tview.NewForm().
		AddInputField("Имя пользователя", "", 20, nil, func(login string) {
			clientInfo.login = login
		}).
		AddPasswordField("Пароль", "", 20, '*', func(pass string) {
			clientInfo.password = pass
		}).
		AddButton("Логин", func() { // Обработка логина
			u, err := client.Login(clientInfo.login, clientInfo.password)
			if err != nil {
				fmt.Println("Game Over")
				time.Sleep(10 * time.Second)
				app.Stop()
			}
			clientInfo.login = u.Login()

			pages.SwitchToPage("Menu")

		}).
		AddButton("Регистрация", func() { // Обработка регистрации
			err := client.SignUp(clientInfo.login, clientInfo.password)
			if err != nil {
				fmt.Println("Game Over")
				time.Sleep(10 * time.Second)
				app.Stop()
			}

			u, err := client.Login(clientInfo.login, clientInfo.password)
			if err != nil {
				fmt.Println("Game Over")
				time.Sleep(10 * time.Second)
				app.Stop()
			}

			clientInfo.login = u.Login()

			pages.SwitchToPage("Menu")
		}).
		AddButton("Выход", func() {
			app.Stop()
		})

	pages.AddPage("Menu", greetingText, true, true)
	pages.AddPage("Auth", AuthForm, true, true)

	flex := tview.NewFlex()
	flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(pages, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 1, false)

	app.SetRoot(flex, true)

	return app
}
