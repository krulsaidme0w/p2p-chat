package repository

import (
	"math/big"
	"sync"

	"p2p-messenger/internal/entity"
	"p2p-messenger/pkg/hash"
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

	p.peers[hash.GetSmallHash(peer.PubKey.String())] = peer
}

func (p *PeerRepository) Delete(pubKey *big.Int) {
	p.rwMutex.RLock()
	defer p.rwMutex.RUnlock()

	delete(p.peers, hash.GetSmallHash(pubKey.String()))
}

func (p *PeerRepository) Get(pubKey *big.Int) (*entity.Peer, bool) {
	p.rwMutex.RLock()
	defer p.rwMutex.RUnlock()

	peer, found := p.peers[hash.GetSmallHash(pubKey.String())]
	return peer, found
}

func (p *PeerRepository) GetPeers() []*entity.Peer {
	peersSlice := make([]*entity.Peer, 0, len(p.peers))

	for _, peer := range p.peers {
		peersSlice = append(peersSlice, peer)
	}

	return peersSlice
}
