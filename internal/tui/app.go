package tui

import (
	"context"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/rivo/tview"
)

type GophClient interface {
	SignUp(login, pass string) error
	Login(login, pass string) (models.User, error)
	SaveLogPass(ctx context.Context, lp *models.LogPassData) error
	SaveCard(ctx context.Context, c *models.Card) error
	SaveText(ctx context.Context, t *models.TextData) error
	SaveBinary(ctx context.Context, t *models.BinaryData) error
}

type App struct {
	ctx        context.Context
	client     GophClient         //client API
	ClientInfo userInfo           // client information
	App        *tview.Application //widgets...
	Pages      *tview.Pages
	AuthForm   *tview.Form
	SaveData   *tview.Flex
	ShowData   *tview.Flex
	Info       *tview.TextView
}

func New(ctx context.Context, client GophClient) *App {

	app := App{
		ctx:        ctx,
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
