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

func (p Chat) Pld() []byte {
	pld := []byte{}
	l := byte(len(p.Txt))

	pld = append(pld, 0xFF)
	pld = append(pld, p.R)
	pld = append(pld, p.G)
	pld = append(pld, p.B)
	pld = append(pld, l)
	pld = append(pld, []byte(p.Txt)...)

	return pld
}
