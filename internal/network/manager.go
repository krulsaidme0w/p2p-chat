package network

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	Port               = 25042
	MulticastIP        = "224.0.0.1"
	MulticastFrequency = 1 * time.Second
)

type Manager struct {
	Listener   *Listener
	Discoverer *Discoverer
}

func NewManager() *Manager {
	multicastAddr, err := net.ResolveUDPAddr(
		"udp",
		fmt.Sprintf("%s:%d", MulticastIP, Port))
	if err != nil {
		log.Fatal(err)
	}

	return &Manager{
		Listener:   NewListener(),
		Discoverer: NewDiscoverer(multicastAddr, MulticastFrequency),
	}
}

func (m *Manager) Start() {
	go m.Discoverer.Start()
}
