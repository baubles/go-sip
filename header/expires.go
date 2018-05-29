package header

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Expires struct {
	Expires int64
}

func NewExpires() *Expires {
	return &Expires{}
}

func (e *Expires) WriteTo(w io.Writer) error {
	return writeStringsTo(w, fmt.Sprintf("%d", e.Expires))
}

func (e *Expires) Marshal() []byte {
	buf := new(bytes.Buffer)
	e.WriteTo(buf)
	return buf.Bytes()
}

func (e *Expires) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	expires, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return fmt.Errorf("unmarshal Expires err: %v", err)
	}
	e.Expires = expires
	return nil
}
