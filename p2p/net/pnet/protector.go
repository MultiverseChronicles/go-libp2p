package pnet

import (
	"errors"
	"net"

	ipnet "github.com/MultiverseChronicles/go-libp2p/pnet"
)

// NewProtectedConn creates a new protected connection
func NewProtectedConn(psk ipnet.PSK, conn net.Conn) (net.Conn, error) {
	if len(psk) != 32 {
		return nil, errors.New("expected 32 byte PSK")
	}
	var p [32]byte
	copy(p[:], psk)
	return newPSKConn(&p, conn)
}
