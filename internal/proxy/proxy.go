package proxy

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/andreyaree/terraria-rooms/internal/metrics"
	"github.com/andreyaree/terraria-rooms/internal/terraria"
	"github.com/andreyaree/terraria-rooms/internal/terraria/packets"
)

func Setup(srvAddr, lstAddr string, bls *Blacklist, m *metrics.Metrics) {
	/* Настраиваем прослушивание порта, на который будут приходить подключения от клиентов, и перенаправляем их на сервер */
	lst, err := net.Listen("tcp", lstAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer lst.Close()

	for {
		clientConn, err := lst.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		addr, _, _ := net.SplitHostPort(clientConn.RemoteAddr().String()) // Получаем и разделяем адрес подключения на две части: IP и порт, и форматируем это в строку

		/* Проверяем в чёрном списке ли адрес или нет, если истина, тогда обрываем соединение */
		if bls.GetStatus(addr) {
			writePacket(clientConn, packets.FatalError{
				Txt: "You are blacklisted :<",
			})
			time.Sleep(time.Second)

			clientConn.Close()
			continue
		}

		go handleConnection(clientConn, srvAddr, m)
	}
}

func handleConnection(clientConn net.Conn, srvAddr string, m *metrics.Metrics) {
	/* Обрабатываем подключение клиента, создаём подключение к серверу и перенаправляем данные между ними */
	m.ActiveConnections.Add(1)
	m.AllTimeConnections.Add(1)

	defer func() {
		m.ActiveConnections.Add(-1)
		clientConn.Close()
	}()

	serverConn, err := net.Dial("tcp", srvAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer serverConn.Close()

	go send(serverConn, clientConn, m, true)
	send(clientConn, serverConn, m, false)
}

func writePacket(conn net.Conn, p terraria.Packet) error {
	pkt := terraria.NewPacket(p)
	_, err := conn.Write(pkt)
	if err != nil {
		return err
	}
	return nil
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
