package conn

import (
	"log"
	"net"
	"time"
)

type peer interface {
	Addr() string
}

func Connect(p peer, PeerID [20]byte, InfoHash [20]byte) {
	log.Println("requesting connection")
	conn, err := net.DialTimeout("tcp", p.Addr(), 3*time.Second)
	if err != nil {
		log.Fatalln("error connecting")
	}
	log.Println("connection setup")

	log.Println("Initiating Handshake")
	h := NewHandShake(PeerID, InfoHash)
	h.DoHandShake(conn)
	log.Println("HandShake done")
}
