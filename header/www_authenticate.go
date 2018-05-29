package header

import (
	"bytes"
	"fmt"
	"io"
)

type WWWAuthenticate struct {
	Credential string
	Params     *Params
}

func NewWWWAuthenticate() *WWWAuthenticate {
	return &WWWAuthenticate{
		Params: NewParamsWithQuote(string(Comma), string(Equals), string(DoubleQuote)),
	}
}

func (wwwauth *WWWAuthenticate) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return fmt.Errorf("unmarshal WWWAuthenticate err: bytes can't be blank string")
	}
	parts := bytes.SplitN(b, []byte{Space}, 2)

	cred := string(parts[0])
	params := NewParamsWithQuote(string(Comma), string(Equals), string(DoubleQuote))
	if len(parts) == 2 {
		if err := params.Unmarshal(parts[1]); err != nil {
			return fmt.Errorf("unmarshal WWWAuthenticate err: %v", err)
		}
	}
	wwwauth.Credential = cred
	wwwauth.Params = params
	return nil
}

func (wwwauth *WWWAuthenticate) WriteTo(w io.Writer) error {
	if err := writeStringsTo(w, wwwauth.Credential, " "); err != nil {
		return err
	}

	if err := wwwauth.Params.WriteTo(w); err != nil {
		return err
	}

	return nil
}

func (wwwauth *WWWAuthenticate) Marshal() []byte {
	buf := new(bytes.Buffer)
	wwwauth.WriteTo(buf)
	return buf.Bytes()
}

func (wwwauth *WWWAuthenticate) SetRealm(realm string) {
	wwwauth.Params.Set(ParamNameRealm, []byte(realm))
}

func (wwwauth *WWWAuthenticate) Realm() (string, bool) {
	val, ok := wwwauth.Params.Get(ParamNameRealm)
	return val.String(), ok
}

func (wwwauth *WWWAuthenticate) SetNonce(nonce string) {
	wwwauth.Params.Set(ParamNameNonce, []byte(nonce))
}

func (wwwauth *WWWAuthenticate) Nonce() (string, bool) {
	val, ok := wwwauth.Params.Get(ParamNameNonce)
	return val.String(), ok
}
