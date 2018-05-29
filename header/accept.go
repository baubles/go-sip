package header

import (
	"bytes"
	"io"
	"strings"
)

type Accept struct {
	Methods []string
}

func NewAccept() *Accept {
	return &Accept{
		Methods: []string{},
	}
}

func (a *Accept) WriteTo(w io.Writer) error {
	return writeStringsTo(w, strings.Join(a.Methods, ", "))
}

func (a *Accept) Marshal() []byte {
	buf := new(bytes.Buffer)
	a.WriteTo(buf)
	return buf.Bytes()
}

func (a *Accept) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		a.Methods = []string{}
		return nil
	}
	parts := bytes.Split(b, []byte(","))
	methods := make([]string, len(parts))
	for i, part := range parts {
		methods[i] = string(bytes.TrimSpace(part))
	}
	return nil
}
