package entity

import (
	b "bytes"
	"errors"
	"math/big"
	"strings"
)

const (
	nullByte = "\x00"
)

var (
	BadMulticastMessage = errors.New("ErrorBadMulticastMessage")
)

type MulticastMessage struct {
	MulticastString string
	Name            string
	PubKey          *big.Int
}

func UdpMulticastMessageToPeer(bytes []byte) (*MulticastMessage, error) {
	bytes = b.Trim(bytes, nullByte)
	array := strings.Split(string(bytes), ":")

	if len(array) != 3 {
		return nil, BadMulticastMessage
	}

	pubKey := new(big.Int)
	pubKey, ok := pubKey.SetString(array[2], 10)
	if !ok {
		return nil, BadMulticastMessage
	}

	return &MulticastMessage{
		MulticastString: array[0],
		Name:            array[1],
		PubKey:          pubKey,
	}, nil
}
