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
	Flex       *tview.Flex
	Info       *tview.TextView
	Pages      *tview.Pages
	AuthForm   *tview.Form
}

func New(client GophClient) *App {

	app := App{
		client:     client,
		ClientInfo: userInfo{login: "unknown user"},
		App:        tview.NewApplication(),
		Flex:       tview.NewFlex(),
		Pages:      tview.NewPages(),
		Info:       tview.NewTextView(),
		AuthForm:   tview.NewForm(),
	}

	return &app
}
