package proto

import (
	"github.com/WolframAlph/dh"

	"p2p-messenger/internal/repository"
)

type Proto struct {
	Name  string
	DH    dh.DiffieHellman
	Peers *repository.PeerRepository
	Port  string
}

func NewProto(name string, port string) *Proto {
	return &Proto{
		Name:  name,
		DH:    dh.New(),
		Peers: repository.NewPeerRepository(),
		Port:  port,
	}
}
