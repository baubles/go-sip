package header

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type MaxForwards struct {
	MaxForwards int
}

func NewMaxForwards() *MaxForwards {
	return &MaxForwards{}
}

func (mf *MaxForwards) WriteTo(w io.Writer) error {
	return writeStringsTo(w, fmt.Sprintf("%d", mf.MaxForwards))
}

func (mf *MaxForwards) Marshal() []byte {
	buf := new(bytes.Buffer)
	mf.WriteTo(buf)
	return buf.Bytes()
}

func (mf *MaxForwards) Unmarshal(b []byte) error {
	b = bytes.TrimSpace(b)
	maxForwards, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return fmt.Errorf("unmarshal MaxForwards err: %v", err)
	}
	mf.MaxForwards = int(maxForwards)
	return nil
}

func (mf *MaxForwards) Clone() HeaderValue {
	return &MaxForwards{
		MaxForwards: mf.MaxForwards,
	}
}
