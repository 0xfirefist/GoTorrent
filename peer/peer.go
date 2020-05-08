package peer

import (
	"encoding/binary"
	"net"
	"strconv"
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

func (p *Peer) Addr() string {
	ip := p.IP.String()
	port := strconv.Itoa(int(p.Port))

	return ip + ":" + port
}
