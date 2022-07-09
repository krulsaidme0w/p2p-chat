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
	ErrBadMulticastMessage = errors.New("ErrorBadMulticastMessage")
)

type MulticastMessage struct {
	MulticastString string
	Name            string
	PubKey          *big.Int
	PubKeyStr       string
	Port            string
}

func UDPMulticastMessageToPeer(bytes []byte) (*MulticastMessage, error) {
	bytes = b.Trim(bytes, nullByte)
	array := strings.Split(string(bytes), ":")

	if len(array) != 4 {
		return nil, ErrBadMulticastMessage
	}

	pubKey := new(big.Int)
	pubKey, ok := pubKey.SetString(array[2], 10)
	if !ok {
		return nil, ErrBadMulticastMessage
	}

	return &MulticastMessage{
		MulticastString: array[0],
		Name:            array[1],
		PubKey:          pubKey,
		PubKeyStr:       array[2],
		Port:            array[3],
	}, nil
}
