package header

import "testing"

func TestURI(t *testing.T) {
	b := []byte("sip:foo@bar.com")
	u := NewURI()
	err := u.Unmarshal(b)
	if err != nil {
		t.Errorf("unmarshal err: %v", err)
	}
	if u.Scheme != "sip" {
		t.Errorf("scheme should be sip but receive: %s", u.Scheme)
	}
	if u.User != "foo" {
		t.Errorf("user should be foo but receive: %s", u.User)
	}
	if u.HostPort.Host != "bar.com" {
		t.Errorf("host should be bar.com but receive: %s", u.HostPort.Host)
	}
	if u.HostPort.Port != 0 {
		t.Errorf("host should be 0 but receive: %v", u.HostPort.Port)
	}

	u.HostPort.Port = 5060
	u.HostPort.Host = "127.0.0.1"
	str := string(u.Marshal())
	if str != "sip:foo@127.0.0.1:5060" {
		t.Errorf(" should be \"sip:foo@127.0.0.1:5060\" but receive: %v", str)
	}

	u.Password = "123456"
	str = string(u.Marshal())
	if str != "sip:foo:123456@127.0.0.1:5060" {
		t.Errorf(" should be \"sip:foo:123456@127.0.0.1:5060\" but receive: %v", str)
	}
}
