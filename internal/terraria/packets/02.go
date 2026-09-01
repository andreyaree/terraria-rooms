package packets

/* Фатальная ошибка. Используем, чтобы выводить свои сообщения клиенту при загрузке и т.д., например если он заблокирован в чёрном списке */
type FatalError struct {
	Txt string
}

func (p FatalError) Type() byte {
	return 0x02
}

func (p FatalError) Pld() []byte {
	pld := []byte{}
	l := byte(len(p.Txt))

	pld = append(pld, 0x00)
	pld = append(pld, l)
	pld = append(pld, []byte(p.Txt)...)

	return pld
}
