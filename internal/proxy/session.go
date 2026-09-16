package proxy

import (
	"net"

	"github.com/andreyaree/terraria-rooms/internal/metrics"
)

type Session struct {
	Client  net.Conn
	Server  net.Conn
	Metrics metrics.Metrics
}

func NewSession(clientConn, serverConn net.Conn) *Session {
	s := &Session{
		Client:  clientConn,
		Server:  serverConn,
		Metrics: metrics.Metrics{},
	}

	return s
}
