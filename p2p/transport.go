package p2p

import "net"

// Peer is an interface that represents the remote nodes
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport is anything that handles the communication between the nodes in the network
type Transport interface {
	Dial(string) error
	Addr() string
	ListenAndAccept() error
	Consume() <-chan RPC // this is read-only channel value returning
	Close() error
}
