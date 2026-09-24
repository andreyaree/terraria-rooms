package proxy

import (
	"io"
	"net"
	"time"

	"github.com/andreyaree/terraria-rooms/internal/metrics"
	"github.com/andreyaree/terraria-rooms/internal/terraria/packets"
)

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return
	}
	defer func() {
		if err := listener.Close(); err != nil {
			return
		}
	}()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			continue
		}

		serverConn, err := net.Dial("tcp", s.ServerAddr)
		if err != nil {
			err := clientConn.Close()
			if err != nil {
				return
			}
			continue
		}

		addr, _, err := net.SplitHostPort(clientConn.RemoteAddr().String())
		if err != nil {
			err := clientConn.Close()
			if err != nil {
				return
			}
			err = serverConn.Close()
			if err != nil {
				return
			}
			continue
		}

		if s.Blacklist.Check(addr) {
			err := writePacket(clientConn, packets.FatalError{
				Txt: "You are blacklisted :<",
			})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
			err = clientConn.Close()
			if err != nil {
				return
			}
			err = serverConn.Close()
			if err != nil {
				return
			}
			continue
		}

		session := NewSession(clientConn, serverConn)
		go session.Run()
	}
}

func (s *Session) Run() {
	s.Metrics.ActiveConnections.Add(1)
	s.Metrics.AllTimeConnections.Add(1)
	defer s.Metrics.ActiveConnections.Add(-1)

	done := make(chan struct{}, 2)

	go func() {
		send(s.Server, s.Client, &s.Metrics, true)
		done <- struct{}{}
	}()
	go func() {
		send(s.Client, s.Server, &s.Metrics, false)
		done <- struct{}{}
	}()

	<-done
	err := s.Client.Close()
	if err != nil {
		return
	}
	err = s.Server.Close()
	if err != nil {
		return
	}
	<-done
}

func send(dst io.Writer, src io.Reader, m *metrics.Metrics, incoming bool) {
	buffer := make([]byte, 32*1024) // создаём буфер размером 32 килобайта, в который будут записываться данные, которые мы будем перенаправлять

	for {
		n, err := src.Read(buffer)
		/* если в буфере есть хоть что-то, то начинаем следующие операции: */
		if n > 0 {
			_, err = dst.Write(buffer[:n]) // берём первые байты и пишем их

			if incoming {
				m.Received.Add(int64(n))
			} else {
				m.Sent.Add(int64(n))
			}
		}
		if err != nil {
			return
		}

	}
}
