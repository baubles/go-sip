package sip

import (
	"context"
	"fmt"
	"sync"

	"github.com/baubles/go-xnet"
)

// Server sip server
type Server struct {
	Addr    string
	Network string

	Handler Handler // handler to invoke

	nsrv    xnet.Server
	mux     sync.Mutex
	running int
}

// Shutdown gracefully shuts down the server without interrupting any active connections
func (srv *Server) Shutdown(ctx context.Context) (err error) {
	srv.mux.Lock()
	if srv.running != 1 {
		err = fmt.Errorf("is not running")
	} else {
		srv.running = 2
	}
	srv.mux.Unlock()
	return srv.nsrv.Shutdown(ctx)
}

// Close server
func (srv *Server) Close() (err error) {
	srv.mux.Lock()
	if srv.running != 1 {
		err = fmt.Errorf("is not running")
	} else {
		srv.running = 2
	}
	srv.mux.Unlock()
	if err != nil {
		return err
	}
	return srv.nsrv.Close()
}

// Serve listens on the network address srv.Addr and then calls Serve to handle requests on incoming connections
func (srv *Server) Serve() (err error) {
	return srv.serve()
}

func (srv *Server) serve() (err error) {
	srv.mux.Lock()
	if srv.running != 0 {
		err = fmt.Errorf("is running")
	} else {
		srv.running = 1
	}
	srv.mux.Unlock()

	if err != nil {
		return err
	}

	defer func() {
		srv.mux.Lock()
		srv.running = 0
		srv.mux.Unlock()
	}()

	switch srv.Network {
	case "udp":
		srv.nsrv = xnet.NewUDPServer(srv.Network, srv.Addr, &netProtocal{}, &netHandler{handler: srv.Handler})
	default:
		return fmt.Errorf("network: %s is not support", srv.Network)
	}

	return srv.nsrv.Serve()
}
