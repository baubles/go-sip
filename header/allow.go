package header

import (
	"bytes"
	"io"
	"strings"
)

type Allow struct {
	Methods []string
}

func NewAllow() *Allow {
	return &Allow{
		Methods: []string{},
	}
}

func (a *Allow) WriteTo(w io.Writer) error {
	return writeStringsTo(w, strings.Join(a.Methods, ", "))
}

func (a *Allow) Marshal() []byte {
	buf := new(bytes.Buffer)
	a.WriteTo(buf)
	return buf.Bytes()
}

func (a *Allow) Unmarshal(b []byte) error {
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
	a.Methods = methods
	return nil
}

func (a *Allow) Clone() HeaderValue {
	clone := &Allow{
		Methods: make([]string, len(a.Methods)),
	}
	copy(clone.Methods, a.Methods)
	return clone
}
