package tui

func (a *App) SetInfoRow(text string) {
	// a.InfoRow.Clear(true)
	a.InfoRow.SetText(text)
}
