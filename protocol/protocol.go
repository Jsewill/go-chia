/* Package Protocol implements types and functions for the golang chia-blockchain protocol */
package protocol

const (
	protocol_version string = "0.0.32"
)

type Capability map[string]uint

type Handshake struct {
	network_id, protocol_version, software_version string
	server_port                                    uint16
	node_type                                      uint8
	capabilities                                   []Capability
}

var (
	Capabilities = []Capability{"BASE": 1}
)

/* Handshake when establishing a connection between two servers. */
