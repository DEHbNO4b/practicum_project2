package tui

import (
	"fmt"

	"github.com/rivo/tview"
)

func SetSaveData(app *App) {

	logPass := tview.NewForm()

	card := tview.NewForm()

	text := tview.NewFlex()

	app.setLogPass(logPass)
	app.setCard(card)
	app.setText(text)

	app.SaveData.
		AddItem(logPass, 0, 1, true).
		AddItem(card, 0, 1, false).
		AddItem(text, 0, 2, false)

}

func (a *App) setLogPass(w *tview.Form) {

	lpd := LogPass{}
	w.SetBorder(true).SetTitle("auth data")
	w.AddInputField("login", "", 20, nil, func(login string) { lpd.Login = login }).
		AddInputField("password", "", 20, nil, func(pass string) { lpd.Pass = pass }).
		AddInputField("info", "", 20, nil, func(info string) { lpd.Info = info }).
		AddButton("save", func() {
			data := lpToDomain(lpd)
			err := a.client.SaveLogPass(a.ctx, &data)
			if err != nil {
				fmt.Println(err)
			}
		})
}

func (a *App) setCard(w *tview.Form) {
	cd := Card{}

	w.SetBorder(true).SetTitle("bankCard data")
	w.AddInputField("cards number", "", 20, nil, func(num string) { cd.CardID = num }).
		AddInputField("pass", "", 20, nil, func(pass string) { cd.Pass = pass }).
		AddInputField("date", "", 20, nil, func(date string) { cd.Date = date }).
		AddInputField("info", "", 20, nil, func(info string) { cd.Info = info }).
		AddButton("save", func() {
			data := cdToDomain(cd)
			err := a.client.SaveCard(a.ctx, &data)
			if err != nil {
				fmt.Println(err)
			}
		})

}

func (a *App) setText(w *tview.Flex) {

	// text := TextData{}

	w.SetBorder(true).SetTitle("text")
	w.SetDirection(tview.FlexRow)

	tArea := tview.NewTextArea().SetBorder(true)
	iArea := tview.NewTextArea().SetBorder(true)
	button := tview.NewButton("save")

	w.AddItem(tArea, 20, 5, false)
	w.AddItem(iArea, 20, 5, false)
	w.AddItem(button, 1, 0, false)

}
