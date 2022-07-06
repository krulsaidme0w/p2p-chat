package repository

import (
	"sync"

	"p2p-messenger/internal/entity"
)

type PeerRepository struct {
	RWMutex *sync.RWMutex
	Peers   map[string]*entity.Peer
}
