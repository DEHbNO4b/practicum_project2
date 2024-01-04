package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) SetSaveData() {

	logPass := tview.NewForm()

	card := tview.NewForm()

	text := tview.NewFlex()

	binary := tview.NewFlex()

	a.setLogPass(logPass)
	a.setCard(card)
	a.setText(text)

	a.SaveData.SetBorder(true).SetTitle("save data")
	a.SaveData.
		AddPage("log-pass", logPass, true, true).
		AddPage("card", card, true, false).
		AddPage("text", text, true, false).
		AddPage("binary", binary, true, false)

	a.SaveData.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 { // 'q'
			a.App.Stop()
		} else if event.Rune() == 108 { // 'l'
			a.SaveData.SwitchToPage("log-pass")
		} else if event.Rune() == 99 { // 'c'
			a.SaveData.SwitchToPage("card")
		} else if event.Rune() == 116 { // 't'
			a.SaveData.SwitchToPage("text")
		} else if event.Rune() == 98 { // 'b'
			a.SaveData.SwitchToPage("log-pass")
		}
		return event
	})
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
