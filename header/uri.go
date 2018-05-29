package header

import (
	"bytes"
	"fmt"
	"io"
)

type URI struct {
	Scheme   string
	User     string
	Password string
	HostPort *HostPort
}

func NewURI() *URI {
	return &URI{
		HostPort: NewHostPort(),
	}
}

func (u *URI) Marshal() []byte {
	buf := new(bytes.Buffer)
	u.WriteTo(buf)
	return buf.Bytes()
}

func (u *URI) WriteTo(w io.Writer) error {
	if err := writeStringsTo(w, u.Scheme, string(Colon), u.User); err != nil {
		return err
	}
	if u.Password != "" {
		if err := writeStringsTo(w, string(Colon), u.Password); err != nil {
			return err
		}
	}
	if err := writeStringsTo(w, string(At)); err != nil {
		return err
	}
	if err := u.HostPort.WriteTo(w); err != nil {
		return err
	}
	return nil
}

func (u *URI) Unmarshal(b []byte) (err error) {
	b = bytes.TrimSpace(b)
	parts := bytes.SplitN(b, []byte{Colon}, 2)
	if len(parts) != 2 {
		return fmt.Errorf("unmarshal uri err: %s", string(b))
	}

	scheme := string(parts[0])

	parts = bytes.SplitN(parts[1], []byte{At}, 2)
	if len(parts) != 2 {
		return fmt.Errorf("unmarshal uir err: %s", string(b))
	}

	hostPort := NewHostPort()
	if err := hostPort.Unmarshal(parts[1]); err != nil {
		return fmt.Errorf("unmarshal uir err: %v", err)
	}

	parts = bytes.SplitN(parts[0], []byte{Colon}, 2)
	user := string(parts[0])
	password := ""
	if len(parts) > 1 {
		password = string(parts[1])
	}

	u.Scheme = scheme
	u.HostPort = hostPort
	u.User = user
	u.Password = password
	return nil
}
