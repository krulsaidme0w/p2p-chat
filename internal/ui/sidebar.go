package ui

import (
	"github.com/rivo/tview"
)

type Sidebar struct {
	View *tview.List
}

func NewSidebar() *Sidebar {
	view := tview.NewList()
	view.SetTitle("peers").SetBorder(true)

	return &Sidebar{
		View: view,
	}
}
