package tui

import "github.com/rivo/tview"

func SetSaveData(app *App) {

	logPass := tview.NewForm().SetBorder(true).SetTitle("auth data")

	card := tview.NewForm().SetBorder(true).SetTitle("card data")

	text := tview.NewForm().SetBorder(true).SetTitle("text data")

	app.SaveData.
		AddItem(logPass, 0, 1, true).
		AddItem(card, 0, 1, false).
		AddItem(text, 0, 2, false)

}

func setLogPass(app *App, w *tview.Form) {
	w.AddInputField("login", "", 35, nil, func(login string) {})
}
