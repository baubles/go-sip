package sip

import (
	"bufio"
	"bytes"
	"io"
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
	return b
}

// Unpack bytes to sip packet
func (proto *netProtocal) Unpack(b []byte) (pkt xnet.Packet, n int, err error) {
	reader := bufio.NewReader(bytes.NewBuffer(b))
	var headline []byte
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("unpack read package bytes error", err)
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
		log.Println("unpack unmarshal pkt err", err)
	}

	return pkt, len(b), nil
}
