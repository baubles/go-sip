package header

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

type Date struct {
	Date time.Time
}

func NewDate() *Date {
	return &Date{Date: time.Now()}
}

func (d *Date) WriteTo(w io.Writer) error {
	return writeStringsTo(w, d.Date.Format(time.RFC1123))
}

func (d *Date) Marshal() []byte {
	buf := new(bytes.Buffer)
	d.WriteTo(buf)
	return buf.Bytes()
}

func (d *Date) Unmarshal(b []byte) error {
	date, err := time.Parse(time.RFC1123, string(bytes.TrimSpace(b)))
	if err != nil {
		return fmt.Errorf("unmarshal date err: %v", err)
	}
	d.Date = date
	return nil
}

func (d *Date) Clone() HeaderValue {
	return &Date{
		Date: d.Date,
	}
}
