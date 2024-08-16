package p2p

const (
	IncomingMessage = 0x1
	IncomingStream  = 0x2
)

// RPC holds data that is being sent over the each transport between two nodes in network
type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
