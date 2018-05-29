package sip

import (
	"github.com/baubles/go-xnet"
)

// Protocal is a implement of sip protocal base xnet
type Protocal struct {
}

// Pack sip packet to bytes
func (proto *Protocal) Pack(pkt xnet.Packet) []byte {
	return pkt.Marshal()
}

// Unpack bytes to sip packet
func (proto *Protocal) Unpack(b []byte) (pkt xnet.Packet, n int, err error) {
	return nil, len(b), nil
}
