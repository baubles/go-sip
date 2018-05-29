package header

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type HostPort struct {
	Host string
	Port int
}

func NewHostPort() *HostPort {
	return &HostPort{}
}

func (hp *HostPort) Marshal() []byte {
	buf := new(bytes.Buffer)
	hp.WriteTo(buf)
	return buf.Bytes()
}

func (hp *HostPort) WriteTo(w io.Writer) error {
	if hp == nil {
		return nil
	}
	if err := writeStringsTo(w, hp.Host); err != nil {
		return err
	}
	if hp.Port > 0 {
		if err := writeStringsTo(w, fmt.Sprintf(":%d", hp.Port)); err != nil {
			return err
		}
	}
	return nil
}

func (hp *HostPort) Unmarshal(b []byte) (err error) {
	parts := bytes.SplitN(b, []byte(":"), 2)

	host := string(parts[0])
	port := int64(0)

	if len(parts) > 1 {
		port, err = strconv.ParseInt(string(parts[1]), 10, 64)
		if err != nil {
			return fmt.Errorf("parse port err: %v", err)
		}
	}

	hp.Host = host
	hp.Port = int(port)

	return nil
}
