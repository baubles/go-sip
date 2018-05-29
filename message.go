package sip

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/baubles/go-sip/header"
)

// Message sip message
type Message struct {
	Headers map[string]header.HeaderValue
	Body    []byte
}

func (msg *Message) Marshal() []byte {
	sep := []byte{CR, LF}
	buf := new(bytes.Buffer)
	for name, val := range msg.Headers {
		buf.Write([]byte(name))
		buf.Write([]byte{header.Colon, ' '})
		val.WriteTo(buf)
		buf.Write(sep)
	}
	buf.Write(sep)
	buf.Write(msg.Body)
	return buf.Bytes()
}

func (msg *Message) Unmarshal(b []byte) error {
	headers := map[string]header.HeaderValue{}
	reader := bufio.NewReader(bytes.NewBuffer(b))
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		parts := bytes.SplitN(line, []byte{header.Colon}, 2)
		name := string(parts[0])
		var val header.HeaderValue
		switch name {
		case header.NameTo:
			val = header.NewTo()
		case header.NameCSeq:
			val = header.NewCSeq()
		case header.NameVia:
			val = header.NewVia()
		case header.NameCallID:
			val = header.NewCallID()
		case header.NameFrom:
			val = header.NewFrom()
		case header.NameMaxForwards:
			val = header.NewMaxForwards()
		case header.NameContentType:
			val = header.NewString()
		case header.NameContentLength:
			val = header.NewContentLength()
		case header.NameDate:
			val = header.NewDate()
		case header.NameAuthorization:
			val = header.NewAuthorization()
		case header.NameWWWAuthenticate:
			val = header.NewWWWAuthenticate()
		case header.NameContact:
			val = header.NewContact()
		case header.NameAllow:
			val = header.NewAllow()
		case header.NameAccept:
			val = header.NewAccept()
		case header.NameExpires:
			val = header.NewExpires()
		default:
			val = header.NewString()
		}
		if err := val.Unmarshal(parts[1]); err != nil {
			return fmt.Errorf("unmarshal Message err: %v", err)
		}
		headers[name] = val
	}

	msg.Body = b[len(b)-reader.Buffered():]
	msg.Headers = headers
	return nil
}
