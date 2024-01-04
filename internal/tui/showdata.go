package tui

import (
	"fmt"

	"github.com/rivo/tview"
)

func (a *App) SetShowData() {

	textView := tview.NewTextView()

	loadButton := tview.NewButton("load data").SetSelectedFunc(func() {
		data, err := a.client.ShowData(a.ctx)
		if err != nil {
			fmt.Println(err)
		}
		textView.SetText(data.String())
	})
	returnButton := tview.NewButton("return").SetSelectedFunc(func() {
		a.Pages.SwitchToPage("save data")
	})

	nf := tview.NewFlex()
	nf.AddItem(tview.NewBox(), 10, 1, false)
	nf.AddItem(loadButton, 20, 1, false)
	nf.AddItem(tview.NewBox(), 5, 1, false)
	nf.AddItem(returnButton, 20, 1, false)
	nf.AddItem(tview.NewBox(), 10, 1, false)

	a.ShowData.SetDirection(tview.FlexRow)
	a.ShowData.AddItem(nf, 1, 1, true)
	a.ShowData.AddItem(textView, 0, 10, false)

}
