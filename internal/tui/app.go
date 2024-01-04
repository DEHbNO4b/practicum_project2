package tui

import (
	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/rivo/tview"
)

type GophClient interface {
	SignUp(login, pass string) error
	Login(login, pass string) (models.User, error)
}

type App struct {
	client     GophClient         //client API
	ClientInfo userInfo           // client information
	App        *tview.Application //widgets...
	Pages      *tview.Pages
	AuthForm   *tview.Form
	SaveData   *tview.Flex
	ShowData   *tview.Flex
	Info       *tview.TextView
}

func New(client GophClient) *App {

	app := App{
		client:     client,
		ClientInfo: userInfo{login: "unknown user"},
		App:        tview.NewApplication(),
		Pages:      tview.NewPages(),
		SaveData:   tview.NewFlex(),
		ShowData:   tview.NewFlex(),
		Info:       tview.NewTextView(),
		AuthForm:   tview.NewForm(),
	}

	setAuthForm(&app)

	SetSaveData(&app)

	setPages(&app)

	app.Pages.SetBorder(true).SetTitle("Goph_keeper")
	app.App.SetRoot(app.Pages, true)

	return &app

}
