package ui

import (
	"github.com/rivo/tview"
)

type InformationField struct {
	View *tview.TextView
}

func NewInformationField() *InformationField {
	view := tview.NewTextView().
		SetText("♡ " + "https://github.com/krulsaidme0w" + " ♡").
		SetTextAlign(tview.AlignCenter)

	view.SetTitle("krulsaidme0w/p2p-messenger").SetBorder(true)

	return &InformationField{
		View: view,
	}
}
