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

func (msg *Message) To() *header.To {
	val, ok := msg.Headers[header.NameTo]
	if ok {
		return val.(*header.To)
	}
	return nil
}

func (msg *Message) CSeq() *header.CSeq {
	val, ok := msg.Headers[header.NameCSeq]
	if ok {
		return val.(*header.CSeq)
	}
	return nil
}

func (msg *Message) Via() *header.Via {
	val, ok := msg.Headers[header.NameVia]
	if ok {
		return val.(*header.Via)
	}
	return nil
}

func (msg *Message) CallID() *header.CallID {
	val, ok := msg.Headers[header.NameCallID]
	if ok {
		return val.(*header.CallID)
	}
	return nil
}

func (msg *Message) From() *header.From {
	val, ok := msg.Headers[header.NameFrom]
	if ok {
		return val.(*header.From)
	}
	return nil
}

func (msg *Message) MaxForwards() *header.MaxForwards {
	val, ok := msg.Headers[header.NameMaxForwards]
	if ok {
		return val.(*header.MaxForwards)
	}
	return nil
}

func (msg *Message) ContentType() *header.String {
	val, ok := msg.Headers[header.NameContentType]
	if ok {
		return val.(*header.String)
	}
	return nil
}

func (msg *Message) ContentLength() *header.ContentLength {
	val, ok := msg.Headers[header.NameContentLength]
	if ok {
		return val.(*header.ContentLength)
	}
	return nil
}

func (msg *Message) Date() *header.Date {
	val, ok := msg.Headers[header.NameDate]
	if ok {
		return val.(*header.Date)
	}
	return nil
}

func (msg *Message) Authorization() *header.Authorization {
	val, ok := msg.Headers[header.NameAuthorization]
	if ok {
		return val.(*header.Authorization)
	}
	return nil
}

func (msg *Message) WWWAuthenticate() *header.WWWAuthenticate {
	val, ok := msg.Headers[header.NameWWWAuthenticate]
	if ok {
		return val.(*header.WWWAuthenticate)
	}
	return nil
}

func (msg *Message) Contact() *header.Contact {
	val, ok := msg.Headers[header.NameContact]
	if ok {
		return val.(*header.Contact)
	}
	return nil
}

func (msg *Message) Allow() *header.Allow {
	val, ok := msg.Headers[header.NameAllow]
	if ok {
		return val.(*header.Allow)
	}
	return nil
}

func (msg *Message) Accept() *header.Accept {
	val, ok := msg.Headers[header.NameAccept]
	if ok {
		return val.(*header.Accept)
	}
	return nil
}

func (msg *Message) Expires() *header.Expires {
	val, ok := msg.Headers[header.NameExpires]
	if ok {
		return val.(*header.Expires)
	}
	return nil
}
