package proto

import (
	"p2p-messenger/internal/repository"

	"github.com/WolframAlph/dh"
)

type Proto struct {
	Name  string
	DH    dh.DiffieHellman
	Peers *repository.PeerRepository
}

func NewProto(name string) *Proto {
	return &Proto{
		Name: name,
		DH:   dh.New(),
	}
}
