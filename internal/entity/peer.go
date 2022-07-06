package entity

import (
	"fmt"
	"math/big"
	"net"
)

type Peer struct {
	Name   string
	PubKey *big.Int
	Conn   *net.Conn
}

func (p *Peer) String() string {
	return fmt.Sprintf("%s:%s", p.Name, p.PubKey)
}
