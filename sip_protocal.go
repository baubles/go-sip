package sip

import (
	"github.com/baubles/go-xnet"
)

// protocal is a implement of sip protocal base xnet
type protocal struct {
}

// Pack sip packet to bytes
func (proto *protocal) Pack(pkt xnet.Packet) []byte {
	return pkt.Marshal()
}

// Unpack bytes to sip packet
func (proto *protocal) Unpack(b []byte) (pkt xnet.Packet, n int, err error) {
	// reader := bufio.NewReader(bytes.NewBuffer(b))

	return nil, len(b), nil
}
