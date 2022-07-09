package ui

import (
	"github.com/rivo/tview"

	"p2p-messenger/internal/repository"
)

type Sidebar struct {
	View             *tview.List
	peerRepo         *repository.PeerRepository
	currentPeerCount int
}

func NewSidebar(peerRepo *repository.PeerRepository) *Sidebar {
	view := tview.NewList()
	view.SetTitle("peers").SetBorder(true)

	return &Sidebar{
		View:             view,
		peerRepo:         peerRepo,
		currentPeerCount: -1,
	}
}

func (s *Sidebar) Reprint() {
	peersCount := len(s.peerRepo.GetPeers())
	if s.currentPeerCount == peersCount {
		return
	}

	s.currentPeerCount = peersCount

	s.View.Clear()

	for _, peer := range s.peerRepo.GetPeers() {
		s.View.
			AddItem(peer.Name, peer.PubKey.String(), 0, nil)
	}
}
