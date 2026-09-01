package metrics

import (
	"fmt"
	"time"
)

func Show(m *Metrics) {
	/* Выводим в консоль метрики, которые мы собираем в процессе работы сервера. Обновляем каждую секунду */
	for {
		fmt.Printf("\rActive Connections: %d, All-Time Connections: %d, Received: %d MB, Sent: %d MB",
			m.ActiveConnections.Load(),
			m.AllTimeConnections.Load(),
			m.Received.Load()/1024/1024, // Получено. Делим до МБ
			m.Sent.Load()/1024/1024)     // Отправлено. Делим до МБ
		time.Sleep(time.Second) // обновляем каждую секунду
	}
}
