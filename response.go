package sip

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"github.com/baubles/go-sip/header"
)

// Response sip response
type Response struct {
	*Message
	Protocal   *header.Protocal
	StatusCode int
	Reason     string

	request *Request
}

func NewResponse() *Response {
	return &Response{
		Message: &Message{},
	}
}

func (res *Response) Unmarshal(b []byte) error {
	var firstline []byte
	reader := bufio.NewReader(bytes.NewBuffer(b))
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return fmt.Errorf("unmarshal Request err: %v", err)
		}
		if len(line) == 0 {
			continue
		}
		firstline = line
		break
	}

	protocal, statusCode, reason, err := res.parseFirstLine(firstline)
	if err != nil {
		return fmt.Errorf("unmarshal Request err: %v", err)
	}

	if err := res.Message.Unmarshal(b[len(b)-reader.Buffered():]); err != nil {
		return fmt.Errorf("unmarshal Request err: %v", err)
	}

	res.StatusCode = statusCode
	res.Protocal = protocal
	res.Reason = reason
	return nil
}

func (res *Response) parseFirstLine(line []byte) (protocal *header.Protocal, statusCode int, reason string, err error) {
	parts := bytes.SplitN(line, []byte{header.Space}, 3)
	if len(parts) < 3 {
		err = fmt.Errorf("parse first line err")
		return
	}

	protocal = header.NewProtocal()
	if err = protocal.Unmarshal(parts[0]); err != nil {
		err = fmt.Errorf("parse first line err: %v", err)
		return
	}

	code, err1 := strconv.ParseInt(string(parts[1]), 10, 64)
	if err1 != nil {
		err = fmt.Errorf("parse first line err: %v", err1)
		return
	}
	statusCode = int(code)
	reason = string(parts[2])
	return
}

func (res *Response) Marshal() []byte {
	buf := new(bytes.Buffer)
	res.Protocal.WriteTo(buf)
	buf.Write([]byte(fmt.Sprintf(" %d %s", res.StatusCode, res.Reason)))
	buf.Write([]byte{CR, LF})
	buf.Write(res.Message.Marshal())
	return buf.Bytes()
}

func (res *Response) Request() *Request {
	return res.request
}
