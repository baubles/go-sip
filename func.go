package sip

import (
	"math/rand"
	"strings"

	"github.com/baubles/go-sip/header"
	"github.com/satori/go.uuid"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CreateResponseFromRequest(req *Request, status int) (res *Response) {
	res = NewResponseWithStatus(status)
	res.SetFrom(req.From().Clone().(*header.From))
	res.SetCSeq(req.CSeq().Clone().(*header.CSeq))
	via := req.Via().Clone().(*header.Via)
	parts := strings.SplitN(req.RemoteAddr.String(), ":", 2)
	if len(parts) == 2 {
		via.SetRPort(parts[1])
	}
	res.SetVia(via)
	to := req.To().Clone().(*header.To)
	res.SetTo(to)
	res.SetCallID(req.CallID())
	return res
}

func uuidString() string {
	u, _ := uuid.NewV4()
	return strings.Replace("-", "", u.String(), -1)
}
