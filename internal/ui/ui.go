package ui

import (
	"time"

	"github.com/rivo/tview"

	"p2p-messenger/internal/proto"
)

const (
	reprintFrequency = 50 * time.Millisecond
)

type App struct {
	Proto     *proto.Proto
	Chat      *Chat
	Sidebar   *Sidebar
	InfoField *InformationField
	View      *tview.Flex
}

func NewApp(proto *proto.Proto) *App {
	app := &App{
		Proto:     proto,
		Chat:      NewChat(),
		Sidebar:   NewSidebar(proto.Peers),
		InfoField: NewInformationField(),
		View:      tview.NewFlex(),
	}

	app.initView()
	app.run()

	return app
}

func (app *App) initView() {
	app.View.
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(app.InfoField.View, 3, 2, false).
			AddItem(app.Sidebar.View, 0, 1, false), 0, 1, false).
		AddItem(app.Chat.View, 0, 3, false)
}

func (app *App) run() {
	ticker := time.NewTicker(reprintFrequency)

	go func(app *App) {
		for {
			select {
			case <-ticker.C:
				app.Sidebar.Reprint()
				app.View.Blur()
			}
		}
	}(app)
}
