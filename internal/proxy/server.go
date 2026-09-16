package proxy

type Server struct {
	Name       string
	ListenAddr string
	ServerAddr string
	Blacklist  *Blacklist
	IsActive   bool
}

func NewServer(name, listenAddr, serverAddr string, blacklist *Blacklist) *Server {
	/* Создаём новую сессию прокси-сервера, которая будет хранить адреса прослушивания и самого сервера, и т.д. */
	s := &Server{
		Name:       name,
		ListenAddr: listenAddr,
		ServerAddr: serverAddr,
		Blacklist:  blacklist,
		IsActive:   true,
	}

	return s
}
