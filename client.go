package sip

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/baubles/go-sip/header"
	"github.com/baubles/go-xnet"
)

type Client struct {
	uri      *header.URI
	password string
	conn     *net.UDPConn
	protocal xnet.Protocal

	connected    bool
	mux          sync.Mutex
	closed       chan bool
	errch        chan error
	accept       chan interface{}
	wg           sync.WaitGroup
	transactions sync.Map
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
		client.errch = make(chan error, 1)
		go func() {
			client.wg.Add(1)
			client.errch <- client.loopRead()
			client.wg.Done()
			client.close(false)
		}()

		go func() {
			client.wg.Add(1)
			client.runTransactionJanitor()
			client.wg.Done()
		}()
	}
	return err
}

func (client *Client) Close() (err error) {
	return client.close(true)
}

func (client *Client) close(force bool) (err error) {
	client.mux.Lock()
	select {
	case <-client.closed:
		client.mux.Unlock()
		return
	default:
		close(client.closed)
		client.mux.Unlock()
	}

	if force {
		err = client.conn.Close()
	}
	client.transactions.Range(func(key, val interface{}) bool {
		trans := val.(*transaction)
		client.transactions.Delete(key)
		close(trans.ch)
		return true
	})
	client.wg.Wait()
	client.connected = false
	return
}

func (client *Client) loopRead() error {
	for {
		buf := make([]byte, 1500)
		n, _, err := client.conn.ReadFrom(buf)

		if err != nil {
			return err
		}
		pkt, _, err := client.protocal.Unpack(buf[:n])
		if err != nil {
			continue
		}

		switch ins := pkt.(type) {
		case *Response:
			tid := fmt.Sprintf("%s:%d", ins.CallID().ID, ins.CSeq().Seq)
			val, ok := client.transactions.Load(tid)
			if ok {
				trans := val.(*transaction)
				select {
				case trans.ch <- ins:
				default:
					log.Println("can't accept pkt, will be drop")
				}

				if ins.StatusCode >= 200 {
					client.transactions.Delete(tid)
					close(trans.ch)
				} else {
					trans.time = time.Now().UnixNano()
				}

				continue
			}
		}

		select {
		case client.accept <- pkt:
		default:
			log.Println("can't accept pkt, will be drop")
		}
	}
}

func (client *Client) runTransactionJanitor() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		client.transactions.Range(func(key, val interface{}) bool {
			trans := val.(*transaction)
			if time.Now().UnixNano()-trans.time > int64(2*time.Second) {
				client.transactions.Delete(key)
				close(trans.ch)
			}
			return true
		})
		select {
		case <-client.closed:
			return
		case <-ticker.C:
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
	_, err := client.conn.Write(client.protocal.Pack(val))
	return err
}

func (client *Client) Register() (err error) {
	expires := header.NewExpires()
	expires.Expires = 3600

	ch, err := client.Request(MethodRegister, 1, "", nil, map[string]header.HeaderValue{
		header.NameExpires: expires,
	})
	if err != nil {
		return err
	}

	res, ok := <-ch
	if !ok {
		return fmt.Errorf("timeout")
	}

	if res.StatusCode == 200 {
		return nil
	}

	if res.StatusCode != 401 {
		return fmt.Errorf("register error, with status code %d", res.StatusCode)
	}

	wwwauth := res.WWWAuthenticate()

	auth := header.NewAuthorization()
	auth.SetUsername(client.uri.User)
	realm, _ := wwwauth.Realm()
	auth.SetRealm(realm)
	nonce, _ := wwwauth.Nonce()
	auth.SetNonce(nonce)
	auth.SetAlgorithm("MD5")
	auth.Credential = "Digest"
	uri := string(client.uri.Marshal())
	auth.SetURI(uri)
	respsonse := GenerateAuthorizationResponse(MethodRegister, client.uri.User, realm, client.password, nonce, uri)
	auth.SetResponse(respsonse)

	headers := map[string]header.HeaderValue{
		header.NameAuthorization: auth,
		header.NameExpires:       expires,
	}

	ch, err = client.Request(MethodRegister, 2, "", nil, headers)

	res, ok = <-ch
	if !ok {
		return fmt.Errorf("timeout")
	}

	if res.StatusCode == 200 {
		return nil
	}

	return fmt.Errorf("register error, with status code %d", res.StatusCode)
}

func (client *Client) Request(method string, seq int64, sipaccount string, body []byte, headers map[string]header.HeaderValue) (res <-chan *Response, err error) {
	req := NewRequest()
	req.Method = method

	via := header.NewVia()
	laddr := client.conn.LocalAddr().(*net.UDPAddr)
	via.SentBy.Host = laddr.IP.String()
	via.SentBy.Port = laddr.Port
	via.Protocal.Transport = "UDP"
	req.SetVia(via)
	via.SetBranch(uuidString())
	via.SetRPort("")

	from := header.NewFrom()
	from.URI = client.uri
	from.SetTag(uuidString())
	req.SetFrom(from)

	contact := header.NewContact()
	contact.URI = client.uri
	req.SetContact(contact)

	to := header.NewTo()
	if sipaccount != "" {
		err = to.URI.Unmarshal([]byte(sipaccount))
		if err != nil {
			return nil, err
		}
	} else {
		to.URI = client.uri
	}
	req.SetTo(to)
	req.URI = to.URI

	maxforwards := header.NewMaxForwards()
	maxforwards.MaxForwards = 70
	req.SetMaxForwards(maxforwards)

	cseq := header.NewCSeq()
	cseq.Method = method
	cseq.Seq = seq
	req.SetCSeq(cseq)

	req.Body = body

	// callid = uuidString()
	callID := &header.CallID{
		ID: uuidString(),
	}
	req.SetCallID(callID)

	for key, val := range headers {
		req.Headers[key] = val
	}

	_, err = client.conn.Write(client.protocal.Pack(req))
	if err != nil {
		return nil, err
	}
	tid := fmt.Sprintf("%s:%d", callID.ID, cseq.Seq)
	trans := newTransaction()
	client.transactions.Store(tid, trans)
	return trans.ch, nil
}

type transaction struct {
	time int64
	ch   chan *Response
}

func newTransaction() *transaction {
	return &transaction{
		time: time.Now().UnixNano(),
		ch:   make(chan *Response, 5),
	}
}
