package header

import (
	"bytes"
	"testing"
)

func TestCSeq(t *testing.T) {
	b := []byte("1 MESSAGE")
	cseq := new(CSeq)
	if err := cseq.Unmarshal(b); err != nil {
		t.Errorf("can't unmarshal \"%s\", err: %v", string(b), err)
	}
	if cseq.Seq != 1 || cseq.Method != "MESSAGE" {
		t.Errorf("unmarshal value is error: %v", cseq)
	}

	tmp := cseq.Marshal()
	if bytes.Compare(b, tmp) != 0 {
		t.Errorf("marshal value is error: \"%s\"", string(tmp))
	}
}
