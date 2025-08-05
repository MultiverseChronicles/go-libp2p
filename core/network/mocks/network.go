package mocknetwork

//go:generate sh -c "go run go.uber.org/mock/mockgen -package mocknetwork -destination mock_resource_manager.go github.com/MultiverseChronicles/go-libp2p/network ResourceManager"
//go:generate sh -c "go run go.uber.org/mock/mockgen -package mocknetwork -destination mock_conn_management_scope.go github.com/MultiverseChronicles/go-libp2p/network ConnManagementScope"
//go:generate sh -c "go run go.uber.org/mock/mockgen -package mocknetwork -destination mock_stream_management_scope.go github.com/MultiverseChronicles/go-libp2p/network StreamManagementScope"
//go:generate sh -c "go run go.uber.org/mock/mockgen -package mocknetwork -destination mock_peer_scope.go github.com/MultiverseChronicles/go-libp2p/network PeerScope"
//go:generate sh -c "go run go.uber.org/mock/mockgen -package mocknetwork -destination mock_protocol_scope.go github.com/MultiverseChronicles/go-libp2p/network ProtocolScope"
//go:generate sh -c "go run go.uber.org/mock/mockgen -package mocknetwork -destination mock_resource_scope_span.go github.com/MultiverseChronicles/go-libp2p/network ResourceScopeSpan"
