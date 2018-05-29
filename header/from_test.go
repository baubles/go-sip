package header

import (
	"bytes"
	"testing"
)

func TestFrom(t *testing.T) {
	b := []byte("<sip:foo@192.168.12.155>;tag=J2eTuwna2")
	from := new(From)
	if err := from.Unmarshal(b); err != nil {
		t.Error(err)
	}
	tag, _ := from.Tag()
	if from.URI.User != "foo" || from.URI.HostPort.Host != "192.168.12.155" || tag != "J2eTuwna2" {
		t.Error("unmarshal result error: ", from.URI, from.Params)
	}

	tmp := from.Marshal()
	if bytes.Compare(tmp, b) != 0 {
		t.Errorf("marshal result error: %s", string(tmp))
	}
}
