package entity

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"time"

	"github.com/WolframAlph/dh"
	"github.com/gorilla/websocket"

	"p2p-messenger/internal/crypto"
)

var (
	ErrPeerIsDeleted = errors.New("peer disconnected")
)

type Peer struct {
	Name      string
	PubKey    *big.Int
	PubKeyStr string
	Port      string
	Messages  []*Message
	AddrIP    string
}

func (p *Peer) AddMessage(text, author string) {
	p.Messages = append(p.Messages, &Message{
		Time:   time.Now(),
		Text:   text,
		Author: author,
	})
}

func (p *Peer) SendMessage(pubKey, message string, dh dh.DiffieHellman) error {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", p.AddrIP, p.Port), Path: "/chat"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if c == nil {
		return ErrPeerIsDeleted
	}

	defer c.Close()
	if err != nil {
		return err
	}

	encryptedMessage, err := crypto.EncryptMessage(crypto.GetSecret(p.PubKey, dh), message)
	if err != nil {
		log.Fatal(err)
	}

	return c.WriteMessage(1, []byte(fmt.Sprintf("%s:%s", pubKey, encryptedMessage)))
}
