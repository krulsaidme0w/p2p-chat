package ui

import (
	"github.com/rivo/tview"
)

type App struct {
	Chat      *Chat
	Sidebar   *Sidebar
	InfoField *InformationField
	View      *tview.Flex
}

func NewApp() *App {
	app := &App{
		Chat:      NewChat(),
		Sidebar:   NewSidebar(),
		InfoField: NewInformationField(),
		View:      tview.NewFlex(),
	}

	app.initView()

	return app
}

func (app *App) initView() {
	app.View.
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(app.InfoField.View, 3, 2, false).
			AddItem(app.Sidebar.View, 0, 1, false), 0, 1, false).
		AddItem(app.Chat.View, 0, 3, false)
}
