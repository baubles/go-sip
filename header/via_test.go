package header

import (
	"fmt"
	"testing"
)

func TestVia(t *testing.T) {
	via := &Via{}
	b := []byte("SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513")
	err := via.Unmarshal(b)
	if err != nil {
		t.Errorf("[TestVia] Error parsing via.  Received: " + err.Error())
	}
	if via.Protocal.Name != "SIP" {
		t.Errorf("[TestVia] Error parsing via \"SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513\".  via.Proto should be \"SIP\" but received: " + via.Protocal.Name)
	}
	if via.Protocal.Version != "2.0" {
		t.Errorf("[TestVia] Error parsing via \"SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513\".  via.Version should be \"2.0\" but received: " + via.Protocal.Version)
	}
	if via.Protocal.Transport != "UDP" {
		t.Errorf("[TestVia] Error parsing via \"SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513\".  via.Transport should be \"UDP\" but received: " + via.Protocal.Transport)
	}
	if via.SentBy.Host != "0.0.0.0" {
		t.Errorf("[TestVia] Error parsing via \"SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513\".  Sent by should be \"0.0.0.0\" but received: " + via.SentBy.Host + ".")
	}

	if via.SentBy.Port != 5060 {
		t.Errorf("[TestVia] Error parsing via \"SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513\".  Sent by should be 5060 but received: %d.", via.SentBy.Port)
	}

	if branch, _ := via.Branch(); branch != "z9hG4bK05B1a4c756d527cb513" {
		t.Errorf("[TestVia] Error parsing via \"SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK05B1a4c756d527cb513\".  Sent by should be \"z9hG4bK05B1a4c756d527cb513\" but received: %s.", branch)
	}

	fmt.Printf("%s\n", via.Marshal())
}
