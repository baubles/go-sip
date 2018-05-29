package sip

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestResponse(t *testing.T) {
	b := []byte(`SIP/2.0 200 OK 
Via: SIP/2.0/UDP 192.168.12.155:5060;rport=5060;received=192.168.12.155;branch=z9hG4bK.Wn6tGbdXF 
Call-ID: y7FOeuQ~Vt 
From: <sip:mingo@192.168.12.155>;tag=FWbGI~g-l 
To: <sip:mingo@192.168.12.155>;tag=z9hG4bK.Wn6tGbdXF 
CSeq: 20 MESSAGE 
Content-Length:  0 
 
`)
	res := NewResponse()
	if err := res.Unmarshal(b); err != nil {
		t.Errorf("unmarshal err: %v", err)
	}

	tmp, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("res: ", string(tmp))

	tmp = res.Marshal()
	fmt.Println("marshal: ", string(tmp))
}
