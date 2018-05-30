package sip

import xnet "github.com/baubles/go-xnet"

// A Handler responds to an HTTP request.
type Handler interface {
	OnRequest(conn xnet.Conn, req *Request)
	OnResponse(conn xnet.Conn, res *Response)
}
