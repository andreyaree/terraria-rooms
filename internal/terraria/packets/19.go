package packets

/* Реализовываем определенный тип пакета - $19, тот, что отвечает за чат */
type Chat struct {
	R   byte
	G   byte
	B   byte
	Txt string
}

func (p Chat) Type() byte {
	return 0x19
}

func (p Chat) Payload() []byte {
	payload := []byte{}
	l := byte(len(p.Txt))

	payload = append(payload, 0xFF)
	payload = append(payload, p.R)
	payload = append(payload, p.G)
	payload = append(payload, p.B)
	payload = append(payload, l)
	payload = append(payload, []byte(p.Txt)...)

	return payload
}
