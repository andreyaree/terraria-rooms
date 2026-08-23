package packets

// Реализовываем определенный тип пакета - $19, тот, что отвечает за чат
type ChatPacket struct {
	R    byte
	G    byte
	B    byte
	Text string
}

func (p ChatPacket) Type() byte {
	return 0x19
}

// Формируем пакет
func (p ChatPacket) Payload() []byte {
	payload := []byte{} // Пустой срез байтов, который предстоит заполнить далее:
	textLength := byte(len(p.Text))

	payload = append(payload, 0xFF)
	payload = append(payload, p.R)
	payload = append(payload, p.G)
	payload = append(payload, p.B)
	payload = append(payload, textLength)
	payload = append(payload, []byte(p.Text)...)

	return payload
}
