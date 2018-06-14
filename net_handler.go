package sip

import (
	"log"
	"sync"

	"github.com/baubles/go-xnet"
)

type netHandler struct {
	conns   sync.Map
	handler Handler
}

func (h *netHandler) Connect(conn xnet.Conn) {
	log.Println("connect raddr:", conn.RemoteAddr())
}

func (h *netHandler) Packet(conn xnet.Conn, pkt xnet.Packet) {
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
	log.Println("disconnect raddr:", conn.RemoteAddr())
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
