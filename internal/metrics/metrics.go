package metrics

import "sync/atomic"

// Создаём структуры где будем хранить значения метрик

type Metrics struct {
	// Используем атомик тип, чтобы правильно изменять переменные из нескольких горутин одновременно (когда запущено несколько комнат)
	// Избегаем гонки данных
	ActiveConnections  atomic.Int32
	AllTimeConnections atomic.Int32
	Received           atomic.Int64
	Sent               atomic.Int64
}
