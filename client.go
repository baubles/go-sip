package sip

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/baubles/go-sip/header"
	"github.com/baubles/go-xnet"
)

type Client struct {
	uri      *header.URI
	password string
	conn     *net.UDPConn
	protocal xnet.Protocal

	connected bool
	mux       sync.Mutex
	closed    chan bool
	errch     chan error
	accept    chan interface{}
	wg        sync.WaitGroup
}

var ErrClientClosed = fmt.Errorf("client be closed")
var ErrClientNotConnected = fmt.Errorf("client is not connected")

func NewClient(sipaccount string, password string) (*Client, error) {
	uri := &header.URI{}
	err := uri.Unmarshal([]byte(sipaccount))
	if err != nil {
		return nil, err
	}

	return &Client{
		uri:      uri,
		password: password,
		protocal: newNetProtocal(),
	}, nil
}

func (client *Client) Dial() (err error) {
	raddr, err := net.ResolveUDPAddr("udp", string(client.uri.HostPort.Marshal()))
	if err != nil {
		return err
	}

	client.mux.Lock()
	defer client.mux.Unlock()

	if client.connected {
		fmt.Errorf("client is connected")
	}

	client.conn, err = net.DialUDP("udp", nil, raddr)
	if err == nil {
		client.connected = true
		client.closed = make(chan bool)
		client.accept = make(chan interface{})
		client.errch = make(chan error)
		go func() {
			client.wg.Add(1)
			client.errch <- client.loop()
			client.wg.Done()
			client.close(false)
		}()
	}
	return err
}

func (client *Client) Close() (err error) {
	return client.close(true)
}

func (client *Client) close(force bool) (err error) {
	client.mux.Lock()
	defer client.mux.Unlock()
	select {
	case <-client.closed:
	default:
		close(client.closed)
		if force {
			err = client.conn.Close()
		}
		client.wg.Wait()
		client.connected = false
	}
	return
}

func (client *Client) loop() error {
	for {
		buf := make([]byte, 1600)
		n, _, err := client.conn.ReadFrom(buf)
		if err != nil {
			return err
		}
		pkt, _, err := client.protocal.Unpack(buf[:n])
		if err != nil {
			log.Println("unpack packet error:", err)
		}

		select {
		case client.accept <- pkt:
		}
	}
}

func (client *Client) Accept() (pkt interface{}, err error) {
	errch := client.errch
	select {
	case pkt = <-client.accept:
		return pkt, nil
	case err = <-errch:
		return nil, err
	case <-client.closed:
		return nil, ErrClientClosed
	}
}

func (client *Client) Write(pkt interface{}) error {
	val, ok := pkt.(xnet.Packet)
	if !ok {
		return fmt.Errorf("pkt can't be cast to xnet.Packet")
	}
	_, err := client.conn.Write(val.Marshal())
	return err
}
