package entity

import (
	"math/big"

	"github.com/gorilla/websocket"
)

type Peer struct {
	Name      string
	PubKey    *big.Int
	PubKeyStr string
	Conn      *websocket.Conn
	Port      string
}
