package tui

import "github.com/rivo/tview"

func (app *App) setMainFlex() {

	app.MainFlex.SetDirection(tview.FlexRow)

	app.MainFlex.AddItem(app.InfoRow, 2, 1, false)
	app.MainFlex.AddItem(app.Pages, 30, 10, true)
}
