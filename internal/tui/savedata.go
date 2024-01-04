package tui

import (
	"fmt"

	"github.com/rivo/tview"
)

func SetSaveData(app *App) {

	logPass := tview.NewForm()

	card := tview.NewForm().SetBorder(true).SetTitle("card data")

	text := tview.NewForm().SetBorder(true).SetTitle("text data")

	app.setLogPass(logPass)
	app.SaveData.
		AddItem(logPass, 0, 1, true).
		AddItem(card, 0, 1, false).
		AddItem(text, 0, 2, false)

}

func (a *App) setLogPass(w *tview.Form) {

	lpd := LogPass{}
	w.SetBorder(true).SetTitle("auth data")
	w.AddInputField("login", "", 25, nil, func(login string) { lpd.Login = login }).
		AddInputField("password", "", 25, nil, func(pass string) { lpd.Pass = pass }).
		AddInputField("info", "", 25, nil, func(info string) { lpd.Info = info }).
		AddButton("save", func() {
			data := lpToDomain(lpd)
			err := a.client.SaveLogPass(a.ctx, data)
			if err != nil {
				fmt.Println(err)
			}
		})
}
