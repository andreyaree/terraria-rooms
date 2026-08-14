package config

import (
	"encoding/json"
	"log"
	"os"
)

// Читаем конфиг из файла по заданному пути и выкидываем ошибки, если это невозможно

func Load(filePath string) Config {
	b, err := os.ReadFile(filePath) // b сокращенно bytes
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Создаем переменную где будем хранить это
	var config Config

	// Unmarshal разбирает данные из JSON и преобразовывает в структуры языка
	err = json.Unmarshal(b, &config) // записываем содержимое файла по адресу в config (который является типом структуры Config)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return config
}
