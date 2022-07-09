package network

import (
	"log"
	"net/http"
	"strings"

	"p2p-messenger/internal/proto"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

type Listener struct {
	proto *proto.Proto
	addr  string
}

func NewListener(addr string, proto *proto.Proto) *Listener {
	return &Listener{
		proto: proto,
		addr:  addr,
	}
}

func (l *Listener) chat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		arr := strings.Split(string(message), ":")
		if len(arr) != 2 {
			continue
		}

		pubKey := arr[0]
		messageText := arr[1]

		peer, found := l.proto.Peers.Get(pubKey)
		if !found {
			continue
		}

		peer.AddMessage(messageText, peer.Name)
	}
}

func (l *Listener) Start() {
	http.HandleFunc("/chat", l.chat)
	log.Fatal(http.ListenAndServe(l.addr, nil))
}
