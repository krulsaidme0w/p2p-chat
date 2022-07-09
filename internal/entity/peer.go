package entity

import (
	"fmt"
	"github.com/WolframAlph/dh"
	"log"
	"math/big"
	"p2p-messenger/internal/crypto"
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

func (p *Peer) SendMessage(pubKey, message string, dh dh.DiffieHellman) error {
	encryptedMessage, err := crypto.EncryptMessage(crypto.GetSecret(p.PubKey, dh), message)
	if err != nil {
		log.Fatal(err)
	}

	return p.Conn.WriteMessage(1, []byte(fmt.Sprintf("%s:%s", pubKey, encryptedMessage)))
}
