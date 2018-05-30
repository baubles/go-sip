package sip

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	srv := &Server{
		Network: "udp",
		Addr:    fmt.Sprintf("%s:5060", localIP()),
	}

	go func() {
		time.Sleep(10 * time.Minute)
		err := srv.Shutdown(context.Background())
		if err != nil {
			t.Errorf("server shutdown err; %v", err)
		}
	}()
	err := srv.Serve()
	if err != nil {
		t.Errorf("server serve err; %v", err)
	}

	// laddr, err := net.ResolveUDPAddr("udp4", ":5060")
	// if err != nil {
	// 	panic(err)
	// }

	// 	net.ListenTCP

	// 	conn, err := net.ListenPacket("udp4", "192.168.13.73:5060")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer conn.Close()

	// 	fmt.Println(conn)

	// 	var (
	// 		addr net.Addr
	// 		n    int
	// 	)
	// 	for {
	// 		byts := make([]byte, 1048)
	// 		n, addr, err = conn.ReadFrom(byts)

	// 		fmt.Println(string(byts), n, addr, err)

	// 		_, err = conn.WriteTo([]byte("\r\n"), addr)

	// 		n, err = conn.WriteTo([]byte(`SIP/2.0 401 Unauthorized
	// Via: SIP/2.0/UDP 192.168.10.67:5060;rport=5060;received=192.168.10.67;branch=z9hG4bK1365646377
	// Call-ID: 1343659654@192.168.10.67
	// From: <sip:34020000001320000001@192.168.10.67>;tag=607342157
	// To: <sip:34020000001320000001@192.168.10.67>;tag=z9hG4bK1365646377
	// CSeq: 1 REGISTER
	// WWW-Authenticate: Digest  realm="192.168.13.73",nonce="56f32f4377a4044d",opaque="31169898427c3c55",algorithm=md5
	// Content-Length:  0

	// `), addr)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		fmt.Println(addr, conn.LocalAddr(), n, err)
	// 	}

}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipAddr, ok := addr.(*net.IPNet); ok && !ipAddr.IP.IsLoopback() {
			if ipAddr.IP.To4() != nil {
				return ipAddr.IP.String()
			}
		}
	}
	return ""
}
