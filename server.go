package sip

import (
	"context"
	"fmt"
	"net"

	"github.com/baubles/go-xnet"
)

// Server sip server
type Server struct {
	Addr    string
	Network string

	Handler Handler // handler to invoke

	nsrv   xnet.Server
	conns  []net.Conn
	closed chan struct{}
}

// Shutdown gracefully shuts down the server without interrupting any active connections
func (srv *Server) Shutdown(ctx context.Context) (err error) {
	// TODO
	return nil
}

// Close server
func (srv *Server) Close() (err error) {
	close(srv.closed)
	return srv.nsrv.Close()
}

// Serve listens on the network address srv.Addr and then calls Serve to handle requests on incoming connections
func (srv *Server) Serve() (err error) {
	return srv.serve()
}

func (srv *Server) listen() (err error) {
	// var nsrv xnet.Server
	switch srv.Network {
	case "udp":
		// nsrv := xnet.NewUDPServer(srv.Network, srv.Addr)

	default:
		return fmt.Errorf("network: %s is not support", srv.Network)
	}
	return nil
}

func (srv *Server) serve() (err error) {
	for {
		// var conn net.Conn
		// conn, err = srv.listener.Accept()
		// if err != nil {
		// 	return err
		// }
	}
	return
}
