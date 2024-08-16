package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr:    ":3000",
		HandShakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}

	tr := NewTCPTranport(opts)

	assert.Equal(t, tr.ListenAddr, ":3000")
	assert.Nil(t, tr.ListenAndAccept())
}
