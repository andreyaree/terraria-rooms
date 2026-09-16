package proxy

type Blacklist struct {
	Addrs map[string]bool
}

func NewBlacklist(addrs []string) *Blacklist {
	/* Создаём чёрный список, который будет хранить адреса в виде ключей в мапе, а значением будет булево значение (true или false) */
	blacklist := &Blacklist{
		Addrs: make(map[string]bool),
	}

	for _, addr := range addrs {
		blacklist.Addrs[addr] = true
	}

	return blacklist
}

func (blacklist *Blacklist) Check(addr string) bool {
	return blacklist.Addrs[addr]
}
