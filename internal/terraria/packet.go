package terraria

import (
	"encoding/binary"
)

/* Интерфейс пакета, который должен реализовывать каждый пакет. Состоит из типа пакета и его данных (payload) */
type Packet interface {
	Type() byte
	Pld() []byte
}

func NewPacket(p Packet) []byte {
	/* Создаём новый пакет, который будет содержать тип пакета и его данные, и общую длину пакета */
	pld := p.Pld()
	l := 1 + len(pld)
	pkt := make([]byte, 3+len(pld))

	binary.LittleEndian.PutUint16(pkt[0:2], uint16(l))
	pkt[2] = p.Type()
	copy(pkt[3:], pld)

	return pkt
}
