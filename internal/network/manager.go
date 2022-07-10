package network

import (
	"fmt"
	"log"
	"net"
	"time"

	"p2p-messenger/internal/proto"
)

const (
	MulticastIP        = "224.0.0.1"
	ListenerIP         = "0.0.0.0"
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
		fmt.Sprintf("%s:%s", MulticastIP, proto.Port))
	if err != nil {
		log.Fatal(err)
	}

	listenerAddr := fmt.Sprintf("%s:%s", ListenerIP, proto.Port)

	return &Manager{
		Proto:      proto,
		Listener:   NewListener(listenerAddr, proto),
		Discoverer: NewDiscoverer(multicastAddr, MulticastFrequency, proto),
	}
}

func (m *Manager) Start() {
	go m.Listener.Start()
	go m.Discoverer.Start()
}
