package terraria

// Пакет Террарии состоит из его типа и последующих данных
type TerrariaPacket interface {
	Type() byte
	Data() []byte
}
