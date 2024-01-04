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

	info := tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText(`
			Press
				(l) to add a new Log-Pass information 
				(c) to add a new bank card information 
				(t) to add a new text information 
				(b) to add a new binary information 
				(s) to load all saved data 
		
		(q) to quit`)
	info.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 { // 'q'
			a.App.Stop()
		} else if event.Rune() == 108 { // 'l'
			a.SaveData.SwitchToPage("log-pass")
		} else if event.Rune() == 99 { // 'c'
			a.SaveData.SwitchToPage("card")
		} else if event.Rune() == 116 { // 't'
			a.SaveData.SwitchToPage("text")
		} else if event.Rune() == 98 { // 'b'
			a.SaveData.SwitchToPage("binary")
		} else if event.Rune() == 115 { // 's'
			a.Pages.SwitchToPage("show data")
		}
		return event
	})

	// a.SaveData.SetBorder(true).SetTitle("save data")
	a.SaveData.
		AddPage("info", info, true, true).
		AddPage("log-pass", logPass, true, false).
		AddPage("card", card, true, false).
		AddPage("text", text, true, false).
		AddPage("binary", binary, true, false)

}

func (a *App) setLogPass(w *tview.Form) {

	lpd := LogPass{}

	w.SetBorder(true).SetTitle("auth data")
	w.AddInputField("login", "", 30, nil, func(login string) { lpd.Login = login }).
		AddInputField("password", "", 30, nil, func(pass string) { lpd.Pass = pass }).
		AddInputField("info", "", 30, nil, func(info string) { lpd.Info = info }).
		AddButton("save", func() {
			data := lpToDomain(lpd)
			err := a.client.SaveLogPass(a.ctx, &data)
			if err != nil {
				fmt.Println(err)
			}
			a.SaveData.SwitchToPage("info")
		}).
		AddButton("return", func() {
			a.SaveData.SwitchToPage("info")
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
		}).
		AddButton("return", func() {
			a.SaveData.SwitchToPage("info")
		})

}

func (a *App) setText(w *tview.Flex) {

	text := TextData{}

	w.SetBorder(true).SetTitle("text")
	w.SetDirection(tview.FlexRow)

	tArea := tview.NewTextArea()
	tArea.SetBorder(true)
	iArea := tview.NewTextArea()
	iArea.SetBorder(true)

	button := tview.NewButton("save").SetSelectedFunc(func() {
		text.Text = tArea.GetText()
		text.Meta = iArea.GetText()
		t := tdToDomain(text)
		err := a.client.SaveText(a.ctx, &t)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("text data saved")
	})
	retButton := tview.NewButton("return").SetSelectedFunc(func() { a.SaveData.SwitchToPage("info") })

	nf := tview.NewFlex()
	nf.AddItem(tview.NewBox(), 10, 1, false)
	nf.AddItem(button, 20, 1, false)
	nf.AddItem(tview.NewBox(), 5, 1, false)
	nf.AddItem(retButton, 20, 1, false)
	nf.AddItem(tview.NewBox(), 10, 1, false)

	w.AddItem(tArea, 0, 5, true)
	w.AddItem(iArea, 0, 5, false)
	w.AddItem(nf, 1, 1, false)

}

func (a *App) setBinary(w *tview.Flex) {

	bd := BinaryData{}

	w.SetBorder(true).SetTitle("saving binary data")
	w.SetDirection(tview.FlexRow)

	tArea := tview.NewTextArea()
	tArea.SetBorder(true)
	iArea := tview.NewTextArea()
	iArea.SetBorder(true)

	button := tview.NewButton("save").SetSelectedFunc(func() {
		bd.Data = tArea.GetText()
		bd.Info = iArea.GetText()
		data := bdToDomain(bd)
		err := a.client.SaveBinary(a.ctx, &data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("binary data saved")
	})
	retButton := tview.NewButton("return").SetSelectedFunc(func() { a.SaveData.SwitchToPage("info") })

	nf := tview.NewFlex()
	nf.AddItem(tview.NewBox(), 10, 1, false)
	nf.AddItem(button, 20, 1, false)
	nf.AddItem(tview.NewBox(), 5, 1, false)
	nf.AddItem(retButton, 20, 1, false)
	nf.AddItem(tview.NewBox(), 10, 1, false)

	w.AddItem(tArea, 0, 5, true)
	w.AddItem(iArea, 0, 5, false)
	w.AddItem(nf, 1, 1, false)

}
