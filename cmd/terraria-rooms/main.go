package main

import (
	"fmt"

	"github.com/andreyaree/terraria-rooms/internal/config"
	"github.com/andreyaree/terraria-rooms/internal/metrics"
	"github.com/andreyaree/terraria-rooms/internal/proxy"
)

func main() {
	config := config.Load("config.json")

	blacklist := proxy.NewBlacklist(config.Blacklist)
	blacklist.ShowBlacklist()

	m := &metrics.Metrics{}
	go metrics.Show(m)

	for _, room := range config.Rooms {
		fmt.Printf("%s: %s -> %s\n", room.Name, room.ListenAddress, room.ServerAddress)
		go proxy.Setup(room.ServerAddress, room.ListenAddress, blacklist, m)
	}
	select {}
}
