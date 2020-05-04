package torrent

import (
	"bytes"
	"crypto/sha1"
	"log"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type file struct {
	Length int    `bencode:"length"`
	Path   string `bencode:"path"`
}

type info struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Length      int    `bencode:"length,omitempty"`
	Files       []file `bencode:"files,omitempty"`
}

type Torrent struct {
	Announce string `bencode:"announce"`
	Info     info   `bencode:"info"`
}

func (t *Torrent) BuildURL(peerID []byte, port uint16) string {

	u, err := url.Parse(t.Announce)
	if err != nil {
		log.Fatalf("There is some problem with URL => %s", err)
	}

	// setting query parameters
	q := url.Values{
		"info_hash":  t.infohash(),
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Info.Length)},
	}

	u.RawQuery = q.Encode()
	return u.String()
}

func (t *Torrent) infohash() []string {
	return []string{t.Info.hash()}
}

func (i *info) hash() string {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *i)
	if err != nil {
		log.Fatalf("Not able to compute infohash => %s", err)
	}
	h := sha1.Sum(buf.Bytes())
	return string(h[:])
}
