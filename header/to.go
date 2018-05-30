package header

import (
	"bytes"
	"fmt"
	"io"
)

type To struct {
	URI    *URI
	Params *Params
}

func NewTo() *To {
	return &To{
		URI:    NewURI(),
		Params: NewParams(string(Semicolon), string(Equals)),
	}
}

func (t *To) Marshal() []byte {
	buf := new(bytes.Buffer)
	t.WriteTo(buf)
	return buf.Bytes()
}

func (t *To) WriteTo(w io.Writer) error {
	if _, err := w.Write([]byte{LeftBrack}); err != nil {
		return err
	}
	if err := t.URI.WriteTo(w); err != nil {
		return err
	}
	if _, err := w.Write([]byte{RightBrack}); err != nil {
		return err
	}
	if t.Params.Size() > 0 {
		if _, err := w.Write([]byte{Semicolon}); err != nil {
			return err
		}

		if err := t.Params.WriteTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *To) Unmarshal(b []byte) error {
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
	t.URI = uir
	t.Params = params
	return nil
}

func (t *To) Clone() HeaderValue {
	return &To{
		URI:    t.URI.Clone().(*URI),
		Params: t.Params.Clone().(*Params),
	}
}
