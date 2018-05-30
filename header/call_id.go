package header

import (
	"bytes"
	"io"
)

type CallID struct {
	ID string
}

func NewCallID() *CallID {
	return &CallID{}
}

func (c *CallID) WriteTo(w io.Writer) error {
	return writeStringsTo(w, c.ID)
}

func (c *CallID) Marshal() []byte {
	buf := new(bytes.Buffer)
	c.WriteTo(buf)
	return buf.Bytes()
}

func (c *CallID) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	c.ID = string(b)
	return nil
}

func (c *CallID) Clone() HeaderValue {
	return &CallID{
		ID: c.ID,
	}
}
