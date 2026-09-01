package packets

/* Сырой (Raw) пакет, который не имеет определённого типа и данных. Используется для передачи данных в их исходном виде */
type Raw struct{}

func (p Raw) Type() byte {
	return p.Type()
}

func (p Raw) Pld() []byte {
	return p.Pld()
}
