package udp

import (
	"net"
)

func ReadFromUDPConnection(conn *net.UDPConn, udpConnectionBufferSize int) ([]byte, *net.UDPAddr, error) {
	buffer := make([]byte, udpConnectionBufferSize)

	_, src, err := conn.ReadFromUDP(buffer)
	if err != nil {
		return nil, nil, err
	}

	return buffer, src, nil
}
