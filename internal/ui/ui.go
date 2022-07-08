package ui

import (
	"github.com/gdamore/tcell/v2"
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
	UI        *tview.Application
}

func NewApp(proto *proto.Proto) *App {
	app := &App{
		Proto:     proto,
		Chat:      NewChat(),
		Sidebar:   NewSidebar(proto.Peers),
		InfoField: NewInformationField(),
		View:      tview.NewFlex(),
		UI:        tview.NewApplication(),
	}

	app.initView()
	app.initUI()
	app.initBindings()

	app.run()

	return app
}

func (app *App) Run() error {
	return app.UI.SetRoot(app.View, true).SetFocus(app.Sidebar.View).Run()
}

func (app *App) initView() {
	app.View.
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(app.InfoField.View, 3, 2, false).
			AddItem(app.Sidebar.View, 0, 1, false), 0, 1, false).
		AddItem(app.Chat.View, 0, 3, false)
}

func (app *App) initUI() {
	app.UI.SetRoot(app.View, true).SetFocus(app.Sidebar.View)
}

func (app *App) initBindings() {
	app.Sidebar.View.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'l' {
			app.UI.SetFocus(app.Chat.Messages)
		}

		if event.Key() == tcell.KeyEnter {
			if app.Sidebar.View.GetItemCount() > 0 {
				app.renderMessages()
			}
		}

		return event
	})

	app.Chat.Messages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'h':
			app.UI.SetFocus(app.Sidebar.View)
		case 'j':
			app.UI.SetFocus(app.Chat.InputField)
		}

		return event
	})

	app.Chat.InputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyUp {
			app.UI.SetFocus(app.Chat.Messages)
		}

		if event.Key() == tcell.KeyEnter {
			app.Chat.InputField.SetText("")
			app.UI.SetFocus(app.Chat.Messages)
		}

		return event
	})
}

func (app *App) renderMessages() {
	app.Chat.Messages.Clear()

	_, pubKey := app.Sidebar.View.GetItemText(
		app.Sidebar.View.GetCurrentItem())

	peer, _ := app.Proto.Peers.Get(pubKey)

	app.Chat.Messages.SetText(peer.Name)

	app.UI.SetFocus(app.Chat.Messages)
}

func (app *App) run() {
	ticker := time.NewTicker(reprintFrequency)

	go func() {
		for {
			<-ticker.C
			app.UI.QueueUpdateDraw(app.Sidebar.Reprint)
		}
	}()
}
