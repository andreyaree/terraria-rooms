package config

import (
	"encoding/json"
	"log"
	"os"
)

func Generate(path string) {
	/* Образец генерируемого файла минимальной конфигурации, который будет создан при отсутствии существующего файла */
	cfg := Config{
		Rooms: []RoomsConfig{
			{
				Name:          "build",
				ServerAddress: "127.0.0.1:7780",
				ListenAddress: "127.0.0.1:7777",
			},
			{
				Name:          "survival",
				ServerAddress: "127.0.0.1:7781",
				ListenAddress: "127.0.0.1:7778",
			},
		},

		Blacklist: []string{},

		GlobalChat: true,

		Messages: MessagesConfig{
			RoomIsFull:       "The room is full",
			RoomNotFound:     "The room not found",
			ConnectionFailed: "The connection failed",
			Switching:        "Switching...",
			Blacklisted:      "You are blacklisted :<",
		},
	}

	data, err := json.MarshalIndent(cfg, " ", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	err = os.WriteFile(path, data, 0755)
	if err != nil {
		log.Fatalf("Failed to generate default config file: %v", err)
	}
}
