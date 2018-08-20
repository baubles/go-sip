package sip

import (
	"bufio"
	"bytes"
	"log"

	"github.com/baubles/go-xnet"
)

// protocal is a implement of sip protocal base xnet
type netProtocal struct {
}

func newNetProtocal() *netProtocal {
	return &netProtocal{}
}

// Pack sip packet to bytes
func (proto *netProtocal) Pack(pkt xnet.Packet) []byte {
	b := pkt.Marshal()
	logger.Println("pack", string(b))
	return b
}

// Unpack bytes to sip packet
func (proto *netProtocal) Unpack(b []byte) (pkt xnet.Packet, n int, err error) {
	logger.Println("unpack", string(b))
	reader := bufio.NewReader(bytes.NewBuffer(b))
	var headline []byte
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println("unpack read package bytes error", err)
			return nil, len(b), err
		}
		line = bytes.TrimSpace(b)
		if len(line) == 0 {
			continue
		}
		headline = line
		break
	}

	if bytes.HasPrefix(headline, []byte("SIP/2.0")) {
		pkt = NewResponse()
	} else {
		pkt = NewRequest()
	}

	if err := pkt.Unmarshal(b); err != nil {
		logger.Printf("unpack unmarshal pkt err: %v, data=%s", err, b)
		return nil, len(b), err
	}

	return pkt, len(b), nil
}
