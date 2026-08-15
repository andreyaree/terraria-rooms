package proxy

import (
	"io"
	"log"
	"net"

	"github.com/andreyaree/terraria-rooms/internal/metrics"
)

// Устанавливаем подключение и перехватываем его, работая с ним в следующей функции
func Setup(serverAddress, listenAddress string, blacklist *Blacklist, m *metrics.Metrics) {
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		// Получаем и разделяем адрес подключения на две части: IP и порт, и форматируем это в строку
		address, _, _ := net.SplitHostPort(clientConn.RemoteAddr().String())

		// Проверяем в чёрном списке ли адрес или нет, если истина, тогда обрываем соединение
		if blacklist.IsBlacklisted(address) {
			clientConn.Close()
			return
		}

		go handleConnection(clientConn, serverAddress, m)
	}
}

// Работа с подключением, добавляем значения в метрику и "перекачиваем" данные на слушающий порт
func handleConnection(clientConn net.Conn, serverAddress string, m *metrics.Metrics) {
	m.ActiveConnections.Add(1)
	m.AllTimeConnections.Add(1)

	defer func() {
		m.ActiveConnections.Add(-1)
		clientConn.Close()
	}()

	serverConn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Println(err)
		return
	}
	defer serverConn.Close()

	clientConn.RemoteAddr()

	// используем функцию-обёртку для копирования данных из сервера к клиенту на слушающий порт
	go copy(serverConn, clientConn, m, true)
	copy(clientConn, serverConn, m, false)
}

// Функция-обёртка, чтобы мы могли считать метрику Отправлено и Получено
// dst есть destination т.е куда пойдут данные и src соотвественно, откуда пришли данные. Incoming - входящее или нет?
func copy(dst io.Writer, src io.Reader, m *metrics.Metrics, incoming bool) {
	buffer := make([]byte, 32*1024) // буффер для хранения данных при чтения из сети, обозначени 32 КБ, что должно хватить на всё

	for {
		n, err := src.Read(buffer)

		// если в буффре есть хоть что-то, то начинаем следующие операции:
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
