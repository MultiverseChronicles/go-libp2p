package metricshelper

import "github.com/MultiverseChronicles/go-libp2p/network"

func GetDirection(dir network.Direction) string {
	switch dir {
	case network.DirOutbound:
		return "outbound"
	case network.DirInbound:
		return "inbound"
	default:
		return "unknown"
	}
}
