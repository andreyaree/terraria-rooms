package packets

/* Фатальная ошибка. Используем, чтобы выводить свои сообщения клиенту при загрузке и т.д., например если он заблокирован в чёрном списке */
type FatalError struct {
	Txt string
}

func (p FatalError) Type() byte {
	return 0x02
}

func (p FatalError) Payload() []byte {
	payload := []byte{}
	l := byte(len(p.Txt))

	payload = append(payload, 0x00)
	payload = append(payload, l)
	payload = append(payload, []byte(p.Txt)...)

	return payload
}
