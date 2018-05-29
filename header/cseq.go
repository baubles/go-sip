package header

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type CSeq struct {
	Seq    int64
	Method string
}

func NewCSeq() *CSeq {
	return &CSeq{}
}

func (c *CSeq) WriteTo(w io.Writer) error {
	if err := writeStringsTo(w, fmt.Sprintf("%d %s", c.Seq, c.Method)); err != nil {
		return err
	}
	return nil
}

func (c *CSeq) Marshal() []byte {
	buf := new(bytes.Buffer)
	c.WriteTo(buf)
	return buf.Bytes()
}

func (c *CSeq) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	parts := bytes.SplitN(b, []byte{Space}, 2)
	if len(parts) != 2 {
		return fmt.Errorf("unmarshal CSeq err: \"%s\" is invalid", string(b))
	}
	seq, err := strconv.ParseInt(string(parts[0]), 10, 64)
	if err != nil {
		fmt.Errorf("unmarshal CSeq parse seq error: %v", err)
	}
	method := string(parts[1])

	c.Seq = seq
	c.Method = method
	return nil
}
