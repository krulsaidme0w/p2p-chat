package network

import (
	"log"
	"net/http"

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

func chat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			break
		}
	}
}

func (l *Listener) Start() {
	http.HandleFunc("/chat", chat)
	log.Fatal(http.ListenAndServe(l.addr, nil))
}
