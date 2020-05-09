package conn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

type HandShake struct {
	Pstr     string
	InfoHash [20]byte
	PeerID   [20]byte
}

func NewHandShake(PeerID [20]byte, InfoHash [20]byte) *HandShake {
	return &HandShake{
		Pstr:     "BitTorrent protocol",
		InfoHash: InfoHash,
		PeerID:   PeerID,
	}
}

func (h *HandShake) DoHandShake(c net.Conn) {
	// write to the connection
	log.Println("Sending Handshake request")
	c.Write(h.serialize())

	// read from the connection
	log.Println("recieving Handshake request")
	resp, _ := ioutil.ReadAll(c)
	fmt.Println(string(resp))
}

// method to serialize handshake struct
func (h *HandShake) serialize() []byte {
	buf := make([]byte, len(h.Pstr)+49)
	buf[0] = byte(len(h.Pstr))
	curr := 1
	curr += copy(buf[curr:], h.Pstr)
	curr += copy(buf[curr:], make([]byte, 8)) // 8 reserved bytes
	curr += copy(buf[curr:], h.InfoHash[:])
	curr += copy(buf[curr:], h.PeerID[:])
	return buf
}
