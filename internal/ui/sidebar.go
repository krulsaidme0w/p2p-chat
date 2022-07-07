package ui

import (
	"github.com/rivo/tview"

	"p2p-messenger/internal/repository"
)

type Sidebar struct {
	View     *tview.List
	PeerRepo *repository.PeerRepository
}

func NewSidebar(peerRepo *repository.PeerRepository) *Sidebar {
	view := tview.NewList()
	view.SetTitle("peers").SetBorder(true)

	return &Sidebar{
		View:     view,
		PeerRepo: peerRepo,
	}
}

func (s *Sidebar) Reprint() {
	s.View.Clear()

	for _, peer := range s.PeerRepo.GetPeers() {
		s.View.
			AddItem(peer.Name, peer.UDPAddr.String(), 0, nil)
	}
}
