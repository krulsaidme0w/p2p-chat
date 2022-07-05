package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Chat struct {
	View       *tview.Flex
	InputField *tview.InputField
	Messages   *tview.TextView
}

func NewChat() *Chat {
	view := tview.NewFlex().SetDirection(tview.FlexRow)
	view.SetTitle("chat").SetBorder(true)

	messages := tview.NewTextView().SetText("").
		SetDynamicColors(true)

	inputField := tview.NewInputField().SetFieldBackgroundColor(tview.Styles.PrimitiveBackgroundColor).
		SetDoneFunc(func(key tcell.Key) {})
	inputField.SetBorder(true)

	view.
		AddItem(messages, 0, 14, false).
		AddItem(inputField, 0, 1, false)

	return &Chat{
		View:       view,
		InputField: inputField,
		Messages:   messages,
	}
}
