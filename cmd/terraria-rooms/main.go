/* Легковесный прокси-сервер для Terraria, с простой конфигурацией и с нужными функциями. Сделал andreyaree */
package main

import (
	"fmt"

	"github.com/andreyaree/terraria-rooms/internal/config"
	"github.com/andreyaree/terraria-rooms/internal/metrics"
	"github.com/andreyaree/terraria-rooms/internal/proxy"
)

func main() {
	/* Загружаем конфигурацию из файла config.json, создаём чёрный список и метрики,
	выводим их в консоль и запускаем прокси-сервер для каждой комнаты */
	cfg := config.Load("config.json")
	bls := proxy.NewBlacklist(cfg.Blacklist)
	m := &metrics.Metrics{}

	/* Выводим дополнительную информацию в консоль */
	bls.Show()
	go metrics.Show(m)

	for _, room := range cfg.Rooms {
		fmt.Printf("%s: %s -> %s\n", room.Name, room.ListenAddress, room.ServerAddress)
		go proxy.Setup(room.ServerAddress, room.ListenAddress, bls, m) // Запускаем прокси-сервер для каждой комнаты в отдельной горутине
	}
	select {}
}
