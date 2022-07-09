package entity

import (
	"fmt"
	"math/big"
	"time"

	"github.com/gorilla/websocket"
)

type Peer struct {
	Name      string
	PubKey    *big.Int
	PubKeyStr string
	Conn      *websocket.Conn
	Port      string
	Messages  []*Message
}

func (p *Peer) AddMessage(text, author string) {
	p.Messages = append(p.Messages, &Message{
		Time:   time.Now(),
		Text:   text,
		Author: author,
	})
}

func (p *Peer) SendMessage(pubKey, message string) error {
	return p.Conn.WriteMessage(1, []byte(fmt.Sprintf("%s:%s", pubKey, message)))
}
