package network

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	udpConnectionBufferSize = 1024
)

type Discoverer struct {
	Addr               *net.UDPAddr
	MulticastFrequency time.Duration
}

func NewDiscoverer(addr *net.UDPAddr, multicastFrequency time.Duration) *Discoverer {
	return &Discoverer{
		Addr:               addr,
		MulticastFrequency: multicastFrequency,
	}
}

func (d *Discoverer) Start() {
	go d.startMulticasting()
	go d.listenMulticasting()
}

func (d *Discoverer) startMulticasting() {
	conn, err := net.DialUDP("udp", nil, d.Addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err := conn.Write([]byte(fmt.Sprintf("meow:%v:%v", "key to communicate", "port to connect")))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(d.MulticastFrequency)
	}
}

func (d *Discoverer) listenMulticasting() {
	conn, err := net.ListenMulticastUDP("udp", nil, d.Addr)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.SetReadBuffer(udpConnectionBufferSize)
	if err != nil {
		log.Fatal(err)
	}

	for {
		buffer := make([]byte, 0, udpConnectionBufferSize)

		_, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		//work with buffer
		fmt.Println(src.IP, string(buffer))
	}
}
