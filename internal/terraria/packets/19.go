package packets

import "github.com/andreyaree/terraria-rooms/internal/utils"

// Реализовываем определенный тип пакета - $19, тот, что отвечает за чат
type ChatPacket struct {
	Color utils.Color
	Text  string
}

func (p ChatPacket) Type() byte {
	return 0x19
}

// Формируем пакет
func (p ChatPacket) Payload() []byte {
	payload := []byte{} // Пустой срез байтов, который предстоит заполнить далее:
	textLength := byte(len(p.Text))

	payload = append(payload, 255)
	payload = append(payload, p.Color.R)
	payload = append(payload, p.Color.G)
	payload = append(payload, p.Color.B)
	payload = append(payload, textLength)
	payload = append(payload, []byte(p.Text)...)

	return payload
}
