package header

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type ContentLength struct {
	ContentLength int64
}

func NewContentLength() *ContentLength {
	return &ContentLength{}
}

func (cl *ContentLength) WriteTo(w io.Writer) error {
	return writeStringsTo(w, fmt.Sprintf("%d", cl.ContentLength))
}

func (cl *ContentLength) Marshal() []byte {
	buf := new(bytes.Buffer)
	cl.WriteTo(buf)
	return buf.Bytes()
}

func (cl *ContentLength) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	contentlength, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return fmt.Errorf("unmarshal Expires err: %v", err)
	}
	cl.ContentLength = contentlength
	return nil
}

func (cl *ContentLength) Clone() HeaderValue {
	return &ContentLength{
		ContentLength: cl.ContentLength,
	}
}
