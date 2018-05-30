package header

import (
	"bytes"
	"fmt"
	"io"
)

type From struct {
	URI    *URI
	Params *Params
}

func NewFrom() *From {
	return &From{URI: NewURI(), Params: NewParams(string(Semicolon), string(Equals))}
}

func (f *From) Tag() (tag string, ok bool) {
	val, ok := f.Params.Get(ParamNameTag)
	return val.String(), ok
}

func (f *From) SetTag(tag string) {
	f.Params.Set(ParamNameTag, StringValue(tag))
}

func (f *From) Marshal() []byte {
	buf := new(bytes.Buffer)
	f.WriteTo(buf)
	return buf.Bytes()
}

func (f *From) WriteTo(w io.Writer) error {
	if _, err := w.Write([]byte{LeftBrack}); err != nil {
		return err
	}
	if err := f.URI.WriteTo(w); err != nil {
		return err
	}
	if _, err := w.Write([]byte{RightBrack}); err != nil {
		return err
	}

	if f.Params.Size() > 0 {
		if _, err := w.Write([]byte{Semicolon}); err != nil {
			return err
		}

		if err := f.Params.WriteTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (f *From) Unmarshal(b []byte) error {
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
		return fmt.Errorf("umarshal From err: %v", err)
	}
	params := NewParams(string(Semicolon), string(Equals))
	if len(parts) > 1 {
		if err := params.Unmarshal(parts[1]); err != nil {
			return fmt.Errorf("umarshal From err: %v", err)
		}
	}
	f.URI = uir
	f.Params = params
	return nil
}

func (f *From) Clone() HeaderValue {
	return &From{
		URI:    f.URI.Clone().(*URI),
		Params: f.Params.Clone().(*Params),
	}
}
