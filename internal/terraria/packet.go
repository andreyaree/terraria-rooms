package terraria

import (
	"encoding/binary"
)

/* Интерфейс пакета, который должен реализовывать каждый пакет. Состоит из типа пакета и его данных (payload) */
type Packet interface {
	Type() byte
	Payload() []byte
}

func NewPacket(p Packet) []byte {
	/* Создаём новый пакет, который будет содержать тип пакета и его данные, и общую длину пакета */
	payload := p.Payload()
	l := 1 + len(payload)
	packet := make([]byte, 3+len(payload))

	binary.LittleEndian.PutUint16(packet[0:2], uint16(l))
	packet[2] = p.Type()
	copy(packet[3:], payload)

	return packet
}
