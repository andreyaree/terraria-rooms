package terraria

type FatalErrorPacket struct {
	ErrorText string
}

func (p FatalErrorPacket) Type() byte {
	return 2
}

func (p FatalErrorPacket) Data() []byte {
	payload := []byte{}

	payload = append(payload, 0x02)

	payload = append(payload, []byte(p.ErrorText)...)

	length := len(payload)

	data := []byte{
		byte(length),
		byte(length >> 8),
	}

	data = append(data, payload...)

	return data
}
