package header

import (
	"bytes"
	"fmt"
	"io"
)

type Authorization struct {
	Credential string
	Params     *Params
}

func NewAuthorization() *Authorization {
	return &Authorization{
		Params: NewParamsWithQuote(string(Comma), string(Equals), string(DoubleQuote)),
	}
}

func (auth *Authorization) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return fmt.Errorf("unmarshal Authorization err: bytes can't be blank string")
	}
	parts := bytes.SplitN(b, []byte{Space}, 2)

	cred := string(parts[0])
	params := NewParamsWithQuote(string(Comma), string(Equals), string(DoubleQuote))
	if len(parts) == 2 {
		if err := params.Unmarshal(parts[1]); err != nil {
			return fmt.Errorf("unmarshal Authorization err: %v", err)
		}
	}
	auth.Credential = cred
	auth.Params = params
	return nil
}

func (auth *Authorization) WriteTo(w io.Writer) error {
	if err := writeStringsTo(w, auth.Credential, " "); err != nil {
		return err
	}

	if err := auth.Params.WriteTo(w); err != nil {
		return err
	}

	return nil
}

func (auth *Authorization) Marshal() []byte {
	buf := new(bytes.Buffer)
	auth.WriteTo(buf)
	return buf.Bytes()
}

func (auth *Authorization) SetUsername(username string) {
	auth.Params.Set(ParamNameUsername, []byte(username))
}

func (auth *Authorization) Username() (string, bool) {
	val, ok := auth.Params.Get(ParamNameUsername)
	return val.String(), ok
}

func (auth *Authorization) SetRealm(realm string) {
	auth.Params.Set(ParamNameRealm, []byte(realm))
}

func (auth *Authorization) Realm() (string, bool) {
	val, ok := auth.Params.Get(ParamNameRealm)
	return val.String(), ok
}

func (auth *Authorization) SetNonce(nonce string) {
	auth.Params.Set(ParamNameNonce, []byte(nonce))
}

func (auth *Authorization) Nonce() (string, bool) {
	val, ok := auth.Params.Get(ParamNameNonce)
	return val.String(), ok
}

func (auth *Authorization) SetURI(uri string) {
	auth.Params.Set(ParamNameURI, []byte(uri))
}

func (auth *Authorization) URI() (string, bool) {
	val, ok := auth.Params.Get(ParamNameURI)
	return val.String(), ok
}

func (auth *Authorization) SetResponse(response string) {
	auth.Params.Set(ParamNameResponse, []byte(response))
}

func (auth *Authorization) Response() (string, bool) {
	val, ok := auth.Params.Get(ParamNameResponse)
	return val.String(), ok
}

func (auth *Authorization) SetAlgorithm(algorithm string) {
	auth.Params.Set(ParamNameAlgorithm, []byte(algorithm))
}

func (auth *Authorization) Algorithm() (string, bool) {
	val, ok := auth.Params.Get(ParamNameAlgorithm)
	return val.String(), ok
}
