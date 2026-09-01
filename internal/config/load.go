package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Load(path string) Config {
	/* Читаем конфиг из файла по заданному пути */
	b, err := os.ReadFile(path) // b сокращение от bytes
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Config not found. Generating new one...")
			Generate(path)
		} else {
			log.Fatalf("Failed to read config file: %v", err)
		}
	}

	var cfg Config

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return cfg
}
