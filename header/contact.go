package header

import (
	"bytes"
	"fmt"
	"io"
)

type Contact struct {
	URI    *URI
	Params *Params
}

func NewContact() *Contact {
	return &Contact{
		URI:    NewURI(),
		Params: NewParams(string(Semicolon), string(Equals)),
	}
}

func (c *Contact) Clone() HeaderValue {
	return &Contact{
		URI:    c.URI.Clone().(*URI),
		Params: c.Params.Clone().(*Params),
	}
}

func (c *Contact) Tag() (tag string, ok bool) {
	val, ok := c.Params.Get(ParamNameTag)
	return val.String(), ok
}

func (c *Contact) SetTag(tag string) {
	c.Params.Set(ParamNameTag, StringValue(tag))
}

func (c *Contact) Marshal() []byte {
	buf := new(bytes.Buffer)
	c.WriteTo(buf)
	return buf.Bytes()
}

func (c *Contact) WriteTo(w io.Writer) error {
	if _, err := w.Write([]byte{LeftBrack}); err != nil {
		return err
	}
	if err := c.URI.WriteTo(w); err != nil {
		return err
	}
	if _, err := w.Write([]byte{RightBrack}); err != nil {
		return err
	}

	if c.Params.Size() > 0 {
		if _, err := w.Write([]byte{Semicolon}); err != nil {
			return err
		}

		if err := c.Params.WriteTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (c *Contact) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return fmt.Errorf("unmarshal err: bytes is empty")
	}
	parts := bytes.SplitN(b, []byte{Semicolon}, 2)

	u := bytes.TrimSpace(parts[0])
	if u[0] != LeftBrack && u[len(u)-1] != RightBrack {

	} else if u[0] == LeftBrack && u[len(u)-1] == RightBrack {
		u = u[1 : len(u)-1]
	} else {
		return fmt.Errorf("unmarshal err: uri \"%s\" not around with <>", string(u))
	}

	uir := NewURI()
	if err := uir.Unmarshal(u); err != nil {
		return fmt.Errorf("umarshal Contact err: %v", err)
	}
	params := NewParams(string(Semicolon), string(Equals))
	if len(parts) > 1 {
		if err := params.Unmarshal(parts[1]); err != nil {
			return fmt.Errorf("umarshal Contact err: %v", err)
		}
	}
	c.URI = uir
	c.Params = params
	return nil
}
