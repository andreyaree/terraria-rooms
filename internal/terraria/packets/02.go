package packets

type FatalErrorPacket struct {
	Text string
}

func (p FatalErrorPacket) Type() byte {
	return 0x02
}

func (p FatalErrorPacket) Payload() []byte {
	payload := []byte{}
	textLength := byte(len(p.Text))

	payload = append(payload, 0x00)
	payload = append(payload, textLength)
	payload = append(payload, []byte(p.Text)...)

	return payload
}
