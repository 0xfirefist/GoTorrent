package peer

import (
	"log"
)

func Unmarshal(stringPeers string) []*Peer {
	const size = 6
	if len(stringPeers)%6 != 0 {
		log.Fatalf("peers info incorrect")
	}

	noPeers := len(stringPeers) / 6
	peers := []*Peer{}

	bytePeers := []byte(stringPeers)

	for i := 0; i < noPeers; i++ {
		index := i * 6
		peers = append(peers, NewPeer(bytePeers[index:index+6]))
	}

	return peers
}
