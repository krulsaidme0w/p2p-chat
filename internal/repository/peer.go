package repository

import (
	"sort"
	"sync"

	"p2p-messenger/internal/entity"
)

type PeerRepository struct {
	rwMutex *sync.RWMutex
	peers   map[string]*entity.Peer
}

func NewPeerRepository() *PeerRepository {
	return &PeerRepository{
		rwMutex: &sync.RWMutex{},
		peers:   make(map[string]*entity.Peer),
	}
}

func (p *PeerRepository) Add(peer *entity.Peer) {
	p.rwMutex.Lock()
	defer p.rwMutex.Unlock()

	_, found := p.peers[peer.PubKeyStr]
	if !found {
		p.peers[peer.PubKeyStr] = peer
	}
}

func (p *PeerRepository) Delete(pubKey string) {
	p.rwMutex.RLock()
	defer p.rwMutex.RUnlock()

	delete(p.peers, pubKey)
}

func (p *PeerRepository) Get(pubKey string) (*entity.Peer, bool) {
	p.rwMutex.RLock()
	defer p.rwMutex.RUnlock()

	peer, found := p.peers[pubKey]
	return peer, found
}

func (p *PeerRepository) GetPeers() []*entity.Peer {
	peersSlice := make([]*entity.Peer, 0, len(p.peers))

	for _, peer := range p.peers {
		peersSlice = append(peersSlice, peer)
	}

	sort.Slice(peersSlice, func(i, j int) bool {
		return peersSlice[i].Name < peersSlice[j].Name
	})

	return peersSlice
}
