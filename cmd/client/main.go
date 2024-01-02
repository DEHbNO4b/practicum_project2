package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	login string
}

var (
	app      = tview.NewApplication()
	contacts []Contact
)

func main() {
	text := tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(q) to quit")

	// box := tview.NewBox().SetBorder(true).SetTitle("Goph_Keeper")

	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	// //create context
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// //create logger
	// // log := setupLogger(cfg.Env)

	// //read config
	// cfg := config.MustLoadClientCfg()

	// //create client
	// client, err := client.New(ctx, cfg)
	// if err != nil {
	// 	panic(err)
	// }

	// switch cfg.Flags.LaunchMode {
	// case "l":
	// 	_, err := client.Login()
	// 	if err != nil {
	// 		return
	// 	}

	// case "r":
	// 	_, err := client.Registert()
	// 	if err != nil {
	// 		return
	// 	}
	// }
}
