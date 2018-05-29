package sip

// Packet is sip net packet
type Packet struct {
}

// Marshal return bytes of packet
func (pkt *Packet) Marshal() []byte {
	return []byte{}
}

func (pkt *Packet) Unmarshal(b []byte) error {
	return nil
}

type HeaderValue interface {
	Bytes()
	String()
}

type Header map[string]HeaderValue
