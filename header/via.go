package header

import (
	"bytes"
	"fmt"
	"io"
)

// Via is sip via header value
type Via struct {
	Protocal *Protocal
	SentBy   *HostPort
	Params   *Params
}

func NewVia() *Via {
	return &Via{
		Protocal: NewProtocal(),
		SentBy:   NewHostPort(),
		Params:   NewParams(string(Semicolon), string(Equals)),
	}
}

// Marshal to bytes
func (v *Via) Marshal() []byte {
	buf := new(bytes.Buffer)
	v.WriteTo(buf)
	return buf.Bytes()
}

func (v *Via) WriteTo(w io.Writer) error {
	if v == nil {
		return nil
	}
	if err := v.Protocal.WriteTo(w); err != nil {
		return err
	}

	if _, err := w.Write([]byte(" ")); err != nil {
		return err
	}

	if err := v.SentBy.WriteTo(w); err != nil {
		return err
	}

	if _, err := w.Write([]byte{Semicolon}); err != nil {
		return err
	}

	if err := v.Params.WriteTo(w); err != nil {
		return err
	}
	return nil
}

func (v *Via) Unmarshal(b []byte) (err error) {
	b = bytes.TrimSpace(b)
	parts := bytes.Split(b, []byte(" "))

	if len(parts) < 2 {
		return fmt.Errorf("Unmarshal err: data format is invalid")
	}

	protocal := NewProtocal()

	if err = protocal.Unmarshal(parts[0]); err != nil {
		return err
	}

	splits := bytes.SplitN(parts[1], []byte{Semicolon}, 2)

	sentBy := NewHostPort()
	if err = sentBy.Unmarshal(splits[0]); err != nil {
		return err
	}

	params := NewParams(string(Semicolon), string(Equals))
	if len(splits) > 1 {
		if err = params.Unmarshal(splits[1]); err != nil {
			return err
		}
	}

	v.Params = params
	v.Protocal = protocal
	v.SentBy = sentBy
	return nil
}

func (v *Via) SetBranch(branch string) {
	v.Params.Set(ParamNameBranch, Value(branch))
}

func (v *Via) Branch() (branch string, ok bool) {
	val, ok := v.Params.Get(ParamNameBranch)
	return val.String(), ok
}

func (v *Via) SetRPort() {
	v.Params.Set(ParamNameRPort, nil)
}

func (v *Via) RPort() (rport int, ok bool) {
	val, ok := v.Params.Get(ParamNameRPort)
	return int(val.Int()), ok
}

func (v *Via) SetTTL(ttl int64) {
	v.Params.Set(ParamNameTTL, IntValue(ttl))
}

func (v *Via) TTL() (ttl int64, ok bool) {
	val, ok := v.Params.Get(ParamNameTTL)
	return val.Int(), ok
}

func (v *Via) SetReceived(received string) {
	v.Params.Set(ParamNameReceived, StringValue(received))
}

func (v *Via) Received() (received string, ok bool) {
	val, ok := v.Params.Get(ParamNameReceived)
	return val.String(), ok
}

func (v *Via) SetMAddr(maddr string) {
	v.Params.Set(ParamNameMAddr, StringValue(maddr))
}

func (v *Via) MAddr() (maddr string, ok bool) {
	val, ok := v.Params.Get(ParamNameMAddr)
	return val.String(), ok
}
