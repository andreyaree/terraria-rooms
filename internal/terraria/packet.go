package terraria

import "encoding/binary"

// Пакет Террарии состоит из его типа и последующих данных
type TerrariaPacket interface {
	Type() byte
	Payload() []byte
}

func NewPacket(p TerrariaPacket) []byte {
	payload := p.Payload()
	length := 1 + len(payload)
	packet := make([]byte, 3+len(payload))

	binary.LittleEndian.PutUint16(packet[0:2], uint16(length))
	packet[2] = p.Type()
	copy(packet[3:], payload)

	return packet
}
