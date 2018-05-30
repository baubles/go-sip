package header

import (
	"bytes"
	"io"
)

type HeaderValue interface {
	WriteTo(w io.Writer) error
	Marshal() []byte
	Clone() HeaderValue
	Unmarshal(b []byte) error
}

type String struct {
	Value string
}

func NewString() *String {
	return &String{}
}

func (s *String) WriteTo(w io.Writer) error {
	return writeStringsTo(w, s.Value)
}

func (s *String) Marshal() []byte {
	buf := new(bytes.Buffer)
	s.WriteTo(buf)
	return buf.Bytes()
}

func (s *String) Unmarshal(b []byte) error {
	s.Value = string(bytes.TrimSpace(b))
	return nil
}

func (s *String) Clone() HeaderValue {
	return &String{
		Value: s.Value,
	}
}
