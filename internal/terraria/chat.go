package terraria

import "github.com/andreyaree/terraria-rooms/internal/utils"

// Реализовываем определенный тип пакета - $19, тот, что отвечает за чат

type ChatPacket struct {
	PlayerSlot byte
	TextColor  utils.Color
	ChatText   string
}

func (p ChatPacket) Type() byte {
	return 19
}

// Формируем пакет

func (p ChatPacket) Data() []byte {
	data := []byte{} // Пустой срез байтов, который предстоит заполнить далее:

	data = append(data, p.Type())
	data = append(data, p.PlayerSlot)
	data = append(data, p.TextColor.R)
	data = append(data, p.TextColor.G)
	data = append(data, p.TextColor.B)
	data = append(data, []byte(p.ChatText)...)

	return data
}
