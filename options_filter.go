package libp2p

import (
	"github.com/MultiverseChronicles/go-libp2p/connmgr"
	"github.com/MultiverseChronicles/go-libp2p/control"
	"github.com/MultiverseChronicles/go-libp2p/network"
	"github.com/MultiverseChronicles/go-libp2p/peer"

	ma "github.com/multiformats/go-multiaddr"
)

// filtersConnectionGater is an adapter that turns multiaddr.Filter into a
// connmgr.ConnectionGater.
type filtersConnectionGater ma.Filters

var _ connmgr.ConnectionGater = (*filtersConnectionGater)(nil)

func (f *filtersConnectionGater) InterceptAddrDial(_ peer.ID, addr ma.Multiaddr) (allow bool) {
	return !(*ma.Filters)(f).AddrBlocked(addr)
}

func (f *filtersConnectionGater) InterceptPeerDial(_ peer.ID) (allow bool) {
	return true
}

func (f *filtersConnectionGater) InterceptAccept(connAddr network.ConnMultiaddrs) (allow bool) {
	return !(*ma.Filters)(f).AddrBlocked(connAddr.RemoteMultiaddr())
}

func (f *filtersConnectionGater) InterceptSecured(_ network.Direction, _ peer.ID, connAddr network.ConnMultiaddrs) (allow bool) {
	return !(*ma.Filters)(f).AddrBlocked(connAddr.RemoteMultiaddr())
}

func (f *filtersConnectionGater) InterceptUpgraded(_ network.Conn) (allow bool, reason control.DisconnectReason) {
	return true, 0
}
