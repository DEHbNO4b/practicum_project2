package widgets

// import (
// 	"fmt"
// 	"time"

// 	"github.com/rivo/tview"
// )

// func New() *App {

// 	app := App{

// 	}

// 	// greetingText := tview.NewTextView().
// 	// 	SetTextColor(tcell.ColorGreen).
// 	// 	SetText("Hello, " + clientInfo.login)

// 	// app :=

// 	// Создаем форму для ввода данных
// 	app.AuthForm.
// 		AddInputField("Имя пользователя", "", 20, nil, func(login string) {
// 			clientInfo.login = login
// 		}).
// 		AddPasswordField("Пароль", "", 20, '*', func(pass string) {
// 			clientInfo.password = pass
// 		}).
// 		AddButton("Логин", func() { // Обработка логина
// 			u, err := client.Login(clientInfo.login, clientInfo.password)
// 			if err != nil {
// 				fmt.Println("Game Over")
// 				time.Sleep(10 * time.Second)
// 				app.Stop()
// 			}
// 			clientInfo.login = u.Login()

// 			pages.SwitchToPage("Menu")

// 		}).
// 		AddButton("Регистрация", func() { // Обработка регистрации
// 			err := client.SignUp(clientInfo.login, clientInfo.password)
// 			if err != nil {
// 				fmt.Println("Game Over")
// 				time.Sleep(10 * time.Second)
// 				app.Stop()
// 			}

// 			u, err := client.Login(clientInfo.login, clientInfo.password)
// 			if err != nil {
// 				fmt.Println("Game Over")
// 				time.Sleep(10 * time.Second)
// 				app.Stop()
// 			}

// 			clientInfo.login = u.Login()

// 			pages.SwitchToPage("Menu")
// 		}).
// 		AddButton("Выход", func() {
// 			app.Stop()
// 		})

// 	pages.AddPage("Menu", greetingText, true, true)
// 	pages.AddPage("Auth", AuthForm, true, true)

// 	// flex :=
// 		flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(pages, 0, 1, false).
// 			AddItem(tview.NewBox().SetBorder(true), 0, 1, false).
// 			AddItem(tview.NewBox().SetBorder(true), 0, 1, false)

// 	app.SetRoot(flex, true)
// }
