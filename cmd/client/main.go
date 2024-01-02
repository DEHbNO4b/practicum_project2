package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DEHbNO4b/practicum_project2/internal/client"
	"github.com/DEHbNO4b/practicum_project2/internal/config"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GophClient interface {
	SignUp(login, pass string) error
	Login(login, pass string) error
}

func main() {

	//create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//create logger
	// log := setupLogger(cfg.Env)

	//read config
	cfg := config.MustLoadClientCfg()

	//create client
	client, err := client.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	app := app(client)

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
func app(client GophClient) *tview.Application {

	user := user{}

	pages := tview.NewPages()

	text := tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(a) to add a new contact \n(q) to quit")

	app := tview.NewApplication()

	// Создаем форму для ввода данных
	form := tview.NewForm().
		AddInputField("Имя пользователя", "", 20, nil, func(login string) {
			user.login = login
		}).
		AddPasswordField("Пароль", "", 20, '*', func(pass string) {
			user.password = pass
		}).
		AddButton("Логин", func() { // Обработка логина
			err := client.Login(user.login, user.password)
			if err != nil {
				fmt.Println("try eshe raz")
				time.Sleep(10 * time.Second)
				app.Stop()
			}
			pages.SwitchToPage("Menu")
		}).
		AddButton("Регистрация", func() { // Обработка регистрации
			err := client.SignUp(user.login, user.password)
			if err != nil {
				fmt.Println("try eshe raz")
				time.Sleep(10 * time.Second)
				app.Stop()
			}
			pages.SwitchToPage("Menu")
		}).
		AddButton("Выход", func() {
			app.Stop()
		})

	pages.AddPage("Menu", text, true, true)
	pages.AddPage("Auth", form, true, true)

	app.SetRoot(pages, true)

	return app
}
