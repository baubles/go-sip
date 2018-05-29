package sip

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	b := []byte(`MESSAGE sip:mingo@192.168.12.155 SIP/2.0 
Via: SIP/2.0/UDP 192.168.12.155:5060;branch=z9hG4bK.GOyhANbhQ;rport 
From: <sip:mingo@192.168.12.155>;tag=3NEN9xjMD 
To: sip:mingo@192.168.12.155 
CSeq: 20 MESSAGE 
Call-ID: NndF9FUKGR 
Max-Forwards: 70 
Supported: replaces, outbound 
Content-Type: text/plain 
Content-Length: 15 
Date: Tue, 29 May 2018 10:12:16 GMT 
User-Agent: Linphone Desktop/4.1.1 (belle-sip/1.6.3) 
 
111
111


12  
`)
	req := NewRequest()
	if err := req.Unmarshal(b); err != nil {
		t.Errorf("unmarshal err: %v", err)
	}

	// tmp, _ := json.MarshalIndent(req, "", "  ")
	// fmt.Println(string(tmp))
	// fmt.Println(string(req.Body))
	tmp := req.Marshal()
	fmt.Println(string(tmp))
}
