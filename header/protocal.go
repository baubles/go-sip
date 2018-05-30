package header

import (
	"bytes"
	"fmt"
	"io"
)

type Protocal struct {
	Name      string
	Version   string
	Transport string
}

func NewProtocal() *Protocal {
	return &Protocal{
		Name:      "SIP",
		Version:   "2.0",
		Transport: "UDP",
	}
}

func (proto *Protocal) Marshal() []byte {
	buf := new(bytes.Buffer)
	proto.WriteTo(buf)
	return buf.Bytes()
}

func (proto *Protocal) WriteTo(w io.Writer) error {
	if proto == nil {
		return nil
	}
	if err := writeStringsTo(w, proto.Name, string(Slash), proto.Version); err != nil {
		return err
	}
	if proto.Transport != "" {
		if err := writeStringsTo(w, string(Slash), proto.Transport); err != nil {
			return err
		}
	}
	return nil
}

func (proto *Protocal) Unmarshal(b []byte) error {
	parts := bytes.Split(b, []byte{Slash})
	if len(parts) < 2 {
		return fmt.Errorf("unmarshal Protocal error")
	}
	name := string(parts[0])
	version := string(parts[1])
	transport := ""
	if len(parts) == 3 {
		transport = string(parts[2])
	}

	proto.Name = name
	proto.Version = version
	proto.Transport = transport

	return nil
}

func (proto *Protocal) Clone() HeaderValue {
	return &Protocal{
		Name:      proto.Name,
		Version:   proto.Version,
		Transport: proto.Transport,
	}
}
