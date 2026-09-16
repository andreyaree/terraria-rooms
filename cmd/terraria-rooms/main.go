/* Легковесный прокси-сервер для Terraria, с простой конфигурацией и с нужными функциями. Сделал andreyaree */
package main

import (
	"github.com/andreyaree/terraria-rooms/internal/config"
	"github.com/andreyaree/terraria-rooms/internal/proxy"
)

func main() {
	/* Загружаем конфигурацию из файла config.json, создаём чёрный список и метрики,
	выводим их в консоль и запускаем прокси-сервер для каждой комнаты */
	cfg := config.Load("config.json")
	blacklist := proxy.NewBlacklist(cfg.Blacklist)

	for _, room := range cfg.Rooms {
		server := proxy.NewServer(room.Name, room.ListenAddress, room.ServerAddress, blacklist)
		go server.Run()
	}
	select {}
}
