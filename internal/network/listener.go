package network

import (
	"p2p-messenger/internal/proto"
)

type Listener struct {
	Proto *proto.Proto
}

func NewListener(proto *proto.Proto) *Listener {
	return &Listener{
		Proto: proto,
	}
}

func (l *Listener) Start() {

}
