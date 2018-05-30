package sip

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	"github.com/baubles/go-sip/header"
)

// A Request represents an SIP request received by a server or to be sent by a client.
type Request struct {
	*Message
	Method   string
	Protocal *header.Protocal
	URI      *header.URI

	LocalAddr  net.Addr
	RemoteAddr net.Addr

	response *Response
}

func NewRequest() *Request {
	return &Request{
		Message: NewMessage(),
	}
}

func (req *Request) Unmarshal(b []byte) error {
	var firstline []byte
	reader := bufio.NewReader(bytes.NewBuffer(b))
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return fmt.Errorf("unmarshal Request err: %v", err)
		}
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		firstline = line
		break
	}

	method, uri, protocal, err := req.parseFirstLine(firstline)
	if err != nil {
		return fmt.Errorf("unmarshal Request err: %v", err)
	}

	if err := req.Message.Unmarshal(b[len(b)-reader.Buffered():]); err != nil {
		return fmt.Errorf("unmarshal Request err: %v", err)
	}

	req.Method = method
	req.URI = uri
	req.Protocal = protocal
	return nil
}

func (req *Request) parseFirstLine(b []byte) (method string, uri *header.URI, protocal *header.Protocal, err error) {
	parts := bytes.SplitN(b, []byte{header.Space}, 3)
	if len(parts) < 3 {
		err = fmt.Errorf("first line is invalid")
		return
	}

	method = string(parts[0])
	uri = header.NewURI()
	if err = uri.Unmarshal(parts[1]); err != nil {
		err = fmt.Errorf("first line err: %v", err)
		return
	}

	protocal = header.NewProtocal()
	if err = protocal.Unmarshal(parts[2]); err != nil {
		err = fmt.Errorf("first line err: %v", err)
		return
	}
	return
}

func (req *Request) Marshal() []byte {
	buf := new(bytes.Buffer)
	buf.Write([]byte(req.Method))
	buf.Write([]byte{header.Space})
	req.URI.WriteTo(buf)
	buf.Write([]byte{header.Space})
	req.Protocal.WriteTo(buf)
	buf.Write([]byte{CR, LF})
	buf.Write(req.Message.Marshal())
	return buf.Bytes()
}

func (req *Request) Response() *Response {
	return req.response
}
