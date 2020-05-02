package torrent

import (
	"fmt"
)

type file struct {
	Length int    `bencode:"length"`
	Path   string `bencode:"path"`
}

type dictionary struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Length      int    `bencode:"length"`
	Files       []file `bencode:"files"`
}

type torrent struct {
	Announce string     `bencode:"announce"`
	Info     dictionary `bencode:"info"`
}

func (t *torrent) BuildURL() (string, error) {
	fmt.Printf("%s\n", *t)
	return "hello", nil
}
