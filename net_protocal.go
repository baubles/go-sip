package sip

import (
	"bufio"
	"bytes"
	"io"

	"go.uber.org/zap"

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
	return pkt.Marshal()
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
			logger.Error("unpack read package bytes error", zap.Error(err))
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
		logger.Error("unpack unmarshal pkt err", zap.Error(err))
	}

	return pkt, len(b), nil
}
