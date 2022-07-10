package network

import (
	"fmt"
	"log"
	"net"
	"time"

	"p2p-messenger/internal/entity"
	"p2p-messenger/internal/proto"
	"p2p-messenger/pkg/udp"
)

const (
	udpConnectionBufferSize = 1024
	multicastString         = "me0w"
)

type Discoverer struct {
	Addr               *net.UDPAddr
	MulticastFrequency time.Duration
	Proto              *proto.Proto
}

func NewDiscoverer(addr *net.UDPAddr, multicastFrequency time.Duration, proto *proto.Proto) *Discoverer {
	return &Discoverer{
		Addr:               addr,
		MulticastFrequency: multicastFrequency,
		Proto:              proto,
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

	ticker := time.NewTicker(d.MulticastFrequency)
	for {
		<-ticker.C
		_, err := conn.Write([]byte(fmt.Sprintf("%s:%s:%s:%s",
			multicastString,
			d.Proto.Name,
			d.Proto.DH.PublicKey,
			d.Proto.Port)))
		if err != nil {
			log.Fatal(err)
		}
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
		rawBytes, addr, err := udp.ReadFromUDPConnection(conn, udpConnectionBufferSize)
		if err != nil {
			log.Fatal(err)
		}

		message, err := entity.UDPMulticastMessageToPeer(rawBytes)
		if err != nil {
			log.Fatal(err)
		}

		peer := &entity.Peer{
			Name:      message.Name,
			PubKey:    message.PubKey,
			PubKeyStr: message.PubKeyStr,
			Port:      message.Port,
			Messages:  make([]*entity.Message, 0),
			AddrIP:    addr.IP.String(),
		}

		if peer.PubKeyStr != d.Proto.DH.PublicKey.String() {
			d.Proto.Peers.Add(peer)
		}
	}
}
