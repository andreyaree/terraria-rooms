package proxy

import "fmt"

// Создаём структуру, которая содержит таблицу значение - ключ (Адрес: true/false - блокировать или нет)
type Blacklist struct {
	Addresses map[string]bool
}

// Новый чёрный список, принимаем в аргументы массив адресов, мы будем брать из конфига, возращая указатель типа Blacklist
func NewBlacklist(addresses []string) *Blacklist {
	blacklist := &Blacklist{
		Addresses: make(map[string]bool),
	}

	for _, address := range addresses {
		blacklist.Addresses[address] = true // Маркируем как чёрный
	}

	return blacklist
}

// Возращаём адрес из таблицы
func (blacklist *Blacklist) IsBlacklisted(address string) bool {
	return blacklist.Addresses[address]
}

func (blacklist *Blacklist) ShowBlacklist() {
	for address := range blacklist.Addresses {
		fmt.Println("Blacklist contains:")
		fmt.Println(address)
	}
}
