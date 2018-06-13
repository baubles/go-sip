package header

import (
	"bytes"
	"io"
)

const (
	ParamNameNextNonce    = "nextnonce"
	ParamNameTag          = "tag"
	ParamNameUsername     = "username"
	ParamNameURI          = "uri"
	ParamNameDomain       = "domain"
	ParamNameCnonce       = "cnonce"
	ParamNamePassword     = "password"
	ParamNameResponse     = "response"
	ParamNameResponseAuth = "rspauth"
	ParamNameOpaque       = "opaque"
	ParamNameAlgorithm    = "algorithm"
	ParamNameDigest       = "digest"
	ParamNameSignedBy     = "signed-by"
	ParamNameSignature    = "signature"
	ParamNameNonce        = "nonce"
	ParamNameSrand        = "srand"
	ParamNameSnum         = "snum"
	ParamNameTargetName   = "targetname"
	ParamNameNonceCount   = "nc"
	ParamNamePubkey       = "pubkey"
	ParamNameCookie       = "cookie"
	ParamNameRealm        = "realm"
	ParamNameVersion      = "version"
	ParamNameStale        = "stale"
	ParamNameQop          = "qop"
	ParamNameNc           = "nc"
	ParamNamePurpose      = "purpose"
	ParamNameCard         = "card"
	ParamNameInfo         = "info"
	ParamNameAction       = "action"
	ParamNameProxy        = "proxy"
	ParamNameRedirect     = "redirect"
	ParamNameExpires      = "expires"
	ParamNameQ            = "q"
	ParamNameRender       = "render"
	ParamNameSession      = "session"
	ParamNameIcon         = "icon"
	ParamNameAlert        = "alert"
	ParamNameHandling     = "handling"
	ParamNameRequired     = "required"
	ParamNameOptional     = "optional"
	ParamNameEmergency    = "emergency"
	ParamNameUrgent       = "urgent"
	ParamNameNormal       = "normal"
	ParamNameNonUrgent    = "non-urgent"
	ParamNameDuration     = "duration"
	ParamNameBranch       = "branch"
	ParamNameHidden       = "hidden"
	ParamNameReceived     = "received"
	ParamNameMAddr        = "maddr"
	ParamNameTTL          = "ttl"
	ParamNameTransport    = "transport"
	ParamNameText         = "text"
	ParamNameCause        = "cause"
	ParamNameID           = "id"
	ParamNameRPort        = "rport"
	ParamNameToTag        = "to-tag"
	ParamNameFromTag      = "from-tag"
	ParamNameSIPInstance  = "+sip.instance"
	ParamNamePubGruu      = "pub-gruu"
	ParamNameTempGruu     = "temp-gruu"
	ParamNameGruu         = "gruu"
)

type Params struct {
	values            map[string]Value
	ListSeparator     string
	KeyValueSeparator string
	Quote             string
}

func NewParams(listSeparator, keyValueSeparator string) *Params {
	return &Params{
		ListSeparator:     listSeparator,
		KeyValueSeparator: keyValueSeparator,
		values:            map[string]Value{},
	}
}

func NewParamsWithQuote(listSeparator, keyValueSeparator, quote string) *Params {
	return &Params{
		ListSeparator:     listSeparator,
		KeyValueSeparator: keyValueSeparator,
		Quote:             quote,
		values:            map[string]Value{},
	}
}

func (p *Params) Unmarshal(b []byte) error {
	values := map[string]Value{}
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return nil
	}
	parts := bytes.Split(b, []byte(p.ListSeparator))
	for _, part := range parts {
		splits := bytes.SplitN(part, []byte(p.KeyValueSeparator), 2)
		if len(splits) == 2 {
			values[string(bytes.TrimSpace(splits[0]))] = Value(bytes.Trim(bytes.TrimSpace(splits[1]), p.Quote))
		} else {
			values[string(splits[0])] = nil
		}
	}
	p.values = values
	return nil
}
func (p *Params) WriteTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	index := 0
	for k, v := range p.values {
		if index > 0 {
			if err := writeStringsTo(w, p.ListSeparator); err != nil {
				return err
			}
		}
		if v == nil || len(v) == 0 {
			if err := writeStringsTo(w, k); err != nil {
				return err
			}
		} else {
			if err := writeTo(w, []byte(k), []byte(p.KeyValueSeparator), []byte(p.Quote), v, []byte(p.Quote)); err != nil {
				return err
			}
		}
		index++
	}
	return nil
}
func (p *Params) Marshal() []byte {
	buf := &bytes.Buffer{}
	p.WriteTo(buf)
	return buf.Bytes()
}

func (p *Params) Clone() HeaderValue {
	values := map[string]Value{}
	for k, v := range p.values {
		values[k] = v.Clone()
	}
	return &Params{
		values:            values,
		KeyValueSeparator: p.KeyValueSeparator,
		ListSeparator:     p.ListSeparator,
		Quote:             p.Quote,
	}
}

func (p *Params) Size() int {
	return len(p.values)
}

func (p *Params) Get(name string) (val Value, ok bool) {
	val, ok = p.values[name]
	return
}

func (p *Params) Set(name string, val Value) {
	p.values[name] = val
}

func (p *Params) Del(name string) {
	delete(p.values, name)
}
