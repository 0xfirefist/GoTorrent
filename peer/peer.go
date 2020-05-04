package peer

import (
	"encoding/binary"
	"net"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func NewPeer(b []byte) *Peer {
	return &Peer{
		IP:   net.IP(b[0:4]),
		Port: binary.BigEndian.Uint16(b[4:6]),
	}
}
