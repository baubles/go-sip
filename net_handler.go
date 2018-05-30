package sip

import (
	"sync"

	"github.com/baubles/go-xnet"
	"go.uber.org/zap"
)

type netHandler struct {
	conns   sync.Map
	handler Handler
}

func (h *netHandler) Connect(conn xnet.Conn) {
	logger.Debug("connect", zap.String("raddr", conn.RemoteAddr().String()))
}

func (h *netHandler) Packet(conn xnet.Conn, pkt xnet.Packet) {
	// logger.Debug("packet", zap.String("raddr", conn.RemoteAddr().String()), zap.ByteString("packet", pkt.Marshal()))
	if h.handler != nil {
		switch ins := pkt.(type) {
		case *Request:
			ins.RemoteAddr = conn.RemoteAddr()
			ins.LocalAddr = conn.LocalAddr()
			h.handler.OnRequest(conn, ins)
		case *Response:
			h.handler.OnResponse(conn, ins)
		}
	}
}

func (h *netHandler) Disconnect(conn xnet.Conn) {
	logger.Debug("disconnect", zap.String("raddr", conn.RemoteAddr().String()))
}

// func (h *netHandler) createResponse(req *Request) {
// 	res := NewResponse()
// 	res.SetFrom(req.From())
// 	res.SetCSeq(req.CSeq())
// 	res.SetVia(req.Via())
// 	res.SetTo(req.To())
// 	req.response = res
// }

// func (h *netHandler) findRequest(res *Response) {

// }
