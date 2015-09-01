package gotcp

import (
	"net"
)

type Packet interface {
	Serialize() []byte
}

type Protocol interface {
	ReadPacket(conn *Conn) (Packet, error)
}
