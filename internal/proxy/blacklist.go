package proxy

import "fmt"

type Blacklist struct {
	Addrs map[string]bool
}

func NewBlacklist(addrs []string) *Blacklist {
	/* Создаём чёрный список, который будет хранить адреса в виде ключей в мапе, а значением будет булево значение (true или false) */
	bls := &Blacklist{
		Addrs: make(map[string]bool),
	}

	for _, addr := range addrs {
		bls.Addrs[addr] = true
	}

	return bls
}

func (bls *Blacklist) GetStatus(addr string) bool {
	return bls.Addrs[addr]
}

func (bls *Blacklist) Show() {
	for addr := range bls.Addrs {
		fmt.Println("Blacklist contains:")
		fmt.Println(addr)
	}
}
