package header

import (
	"io"
	"strconv"
)

type Value []byte

func (v Value) String() string {
	return string(v)
}

func (v Value) Int() int64 {
	i, _ := strconv.ParseInt(v.String(), 10, 64)
	return i
}

func (v Value) Bool() bool {
	b, _ := strconv.ParseBool(v.String())
	return b
}

func IntValue(i int64) Value {
	return []byte(strconv.FormatInt(i, 10))
}

func StringValue(s string) Value {
	return []byte(s)
}

func BoolValue(b bool) Value {
	return []byte(strconv.FormatBool(b))
}

func (v Value) Clone() Value {
	clone := make(Value, len(v))
	copy(clone, v)
	return clone
}

func writeStringsTo(w io.Writer, ss ...string) error {
	for _, s := range ss {
		if _, err := w.Write([]byte(s)); err != nil {
			return err
		}
	}
	return nil
}

func writeTo(w io.Writer, bs ...[]byte) error {
	for _, b := range bs {
		if _, err := w.Write(b); err != nil {
			return err
		}
	}
	return nil
}
