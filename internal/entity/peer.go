package entity

import (
	"math/big"
	"net"
)

type Peer struct {
	Name    string
	PubKey  *big.Int
	Conn    *net.Conn
	UDPAddr *net.UDPAddr
}
