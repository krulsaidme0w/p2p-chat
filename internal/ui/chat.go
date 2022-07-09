package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"p2p-messenger/internal/entity"
)

const (
	timeFormat = "15:04:05"
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

func (c *Chat) RenderMessages(messages []*entity.Message) {
	text := ""

	for i := 0; i < 100; i++ {
		text += "\n"
	}

	for _, message := range messages {
		text += fmt.Sprintf("%s %s: %s\n",
			formatTime(message),
			formatAuthor(message),
			formatText(message))
	}

	c.Messages.SetText(text[:len(text)-1]).ScrollToEnd()
}

func formatTime(message *entity.Message) string {
	now := message.Time.UTC()
	return fmt.Sprintf("%s%s", "[blue]", now.Format(timeFormat))
}

func formatAuthor(message *entity.Message) string {
	return fmt.Sprintf("%s%s", "[red]", message.Author)
}

func formatText(message *entity.Message) string {
	return fmt.Sprintf("%s%s", "[white]", message.Text)
}
