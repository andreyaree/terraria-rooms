package config

/* Структура конфига, которая будет хранить все настройки программы */
type Config struct {
	Rooms      []RoomsConfig  `json:"rooms"`
	Blacklist  []string       `json:"blacklist"`
	GlobalChat bool           `json:"global_chat_enable"`
	Messages   MessagesConfig `json:"messages"`
}

type RoomsConfig struct {
	Name          string `json:"name"`
	ServerAddress string `json:"target"`
	ListenAddress string `json:"listen"`
}

type MessagesConfig struct {
	RoomIsFull       string `json:"room_is_full"`
	RoomNotFound     string `json:"room_not_found"`
	ConnectionFailed string `json:"connection_failed"`
	Switching        string `json:"switching"`
	Blacklisted      string `json:"blacklisted"`
}
