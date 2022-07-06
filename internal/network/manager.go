package network

import (
	"fmt"
	"log"
	"net"
	"time"

	"p2p-messenger/internal/proto"
)

const (
	Port               = 25042
	MulticastIP        = "224.0.0.1"
	MulticastFrequency = 1 * time.Second
)

type Manager struct {
	Proto      *proto.Proto
	Listener   *Listener
	Discoverer *Discoverer
}

func NewManager(proto *proto.Proto) *Manager {
	multicastAddr, err := net.ResolveUDPAddr(
		"udp",
		fmt.Sprintf("%s:%d", MulticastIP, Port))
	if err != nil {
		log.Fatal(err)
	}

	return &Manager{
		Proto:      proto,
		Listener:   NewListener(proto),
		Discoverer: NewDiscoverer(multicastAddr, MulticastFrequency, proto),
	}
}

func (m *Manager) Start() {
	go m.Discoverer.Start()
}
