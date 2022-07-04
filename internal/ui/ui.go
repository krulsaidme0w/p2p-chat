package ui

import (
	"github.com/bxcodec/faker/v3"
	"github.com/rivo/tview"
)

type App struct {
	tApp              *tview.Application
	tFlex             *tview.Flex
	tInformationField *tview.TextView
	tPeersSideBar     *tview.List
}

func NewApp() *App {
	app := new(App)

	app.initInformationField()
	app.initPeersSideBar()
	app.initFlex()
	app.initApp()

	return app
}

func (app *App) Run() error {
	if err := app.tApp.SetRoot(app.tFlex, true).SetFocus(app.tPeersSideBar).Run(); err != nil {
		return err
	}
	return nil
}

func (app *App) initApp() {
	app.tApp = tview.NewApplication()
}

func (app *App) initFlex() {
	app.tFlex = tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(app.tInformationField, 3, 2, false).
			AddItem(app.tPeersSideBar, 0, 1, false), 0, 1, false).
		AddItem(tview.NewTextView().SetBorder(true).SetTitle("chats"), 0, 3, false)
}

func (app *App) initInformationField() {
	app.tInformationField = tview.NewTextView().SetText("https://github.com/krulsaidme0w").SetTextAlign(tview.AlignCenter)
	app.tInformationField.SetTitle("krulsaidme0w/p2p-messenger").SetBorder(true)
}

func (app *App) initPeersSideBar() {
	app.tPeersSideBar = tview.NewList()
	app.tPeersSideBar.SetTitle("peers").SetBorder(true)

	for i := 0; i < 10; i++ {
		app.tPeersSideBar.AddItem(faker.Name(), faker.Password(), 0, nil)
	}
}
