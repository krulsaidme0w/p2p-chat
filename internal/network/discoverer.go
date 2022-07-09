package network

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/url"
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

	for {
		_, err := conn.Write([]byte(fmt.Sprintf("%s:%s:%s:%s",
			multicastString,
			d.Proto.Name,
			d.Proto.DH.PublicKey,
			d.Proto.Port)))
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
		rawBytes, addr, err := udp.ReadFromUDPConnection(conn, udpConnectionBufferSize)
		if err != nil {
			log.Fatal(err)
		}

		message, err := entity.UdpMulticastMessageToPeer(rawBytes)
		if err != nil {
			log.Fatal(err)
		}

		peer := &entity.Peer{
			Name:      message.Name,
			PubKey:    message.PubKey,
			PubKeyStr: message.PubKeyStr,
			Conn:      nil,
			Port:      message.Port,
			Messages:  make([]*entity.Message, 0),
		}

		go d.connectToPeer(peer, fmt.Sprintf("%s:%s", addr.IP.String(), peer.Port))
	}
}

func (d *Discoverer) connectToPeer(peer *entity.Peer, addr string) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/chat"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	peer.Conn = c

	d.Proto.Peers.Add(peer)
}
