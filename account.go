package sip

import "net"

// A Account represents a sip account
type Account struct {
	Realm    string // Realm. Use "*" to make a credential that can be used to authenticate against any challenges.
	Scheme   string // Scheme (e.g. "digest").
	DataType int    //Type of data (0 for plaintext passwd).
	Data     string // The data, which can be a plaintext password or a hashed digest.
	InOnline bool

	Conn net.Conn
}

// Send send buf to account
func (acc *Account) Send(buf []byte) (n int, err error) {
	return acc.Conn.Write(buf)
}

// SendRequest send sip request to account
func (acc *Account) SendRequest(req *Request) (err error) {
	// _, err = acc.Conn.Write(req.Bytes())
	return err
}
