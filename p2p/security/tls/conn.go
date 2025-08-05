package libp2ptls

import (
	"crypto/tls"

	ci "github.com/MultiverseChronicles/go-libp2p/crypto"
	"github.com/MultiverseChronicles/go-libp2p/network"
	"github.com/MultiverseChronicles/go-libp2p/peer"
	"github.com/MultiverseChronicles/go-libp2p/sec"
)

type conn struct {
	*tls.Conn

	localPeer       peer.ID
	remotePeer      peer.ID
	remotePubKey    ci.PubKey
	connectionState network.ConnectionState
}

var _ sec.SecureConn = &conn{}

func (c *conn) LocalPeer() peer.ID {
	return c.localPeer
}

func (c *conn) RemotePeer() peer.ID {
	return c.remotePeer
}

func (c *conn) RemotePublicKey() ci.PubKey {
	return c.remotePubKey
}

func (c *conn) ConnState() network.ConnectionState {
	return c.connectionState
}
