package proxy

import (
	"net"

	"github.com/andreyaree/terraria-rooms/internal/terraria"
)

func writePacket(conn net.Conn, p terraria.Packet) error {
	packet := terraria.NewPacket(p)
	_, err := conn.Write(packet)
	if err != nil {
		return err
	}
	return nil
}
