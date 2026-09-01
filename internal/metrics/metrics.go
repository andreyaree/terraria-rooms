package metrics

import "sync/atomic"

/*
Структура для хранения метрик, которые мы будем собирать в процессе работы прокси-сервера.
Метрики включают количество активных соединений, общее количество соединений за всё время работы,
количество полученных и отправленных байтов
*/
type Metrics struct {
	ActiveConnections  atomic.Int32
	AllTimeConnections atomic.Int32
	Received           atomic.Int64
	Sent               atomic.Int64
}
