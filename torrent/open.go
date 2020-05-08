package torrent

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/jackpal/bencode-go"
)

func Open(path string) *Torrent {

	// file read
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("file not found - %s", path)
	}

	// unmarshal
	log.Println("Checking file contents")
	t := Torrent{}
	err = bencode.Unmarshal(bytes.NewReader(fileContent), &t)
	if err != nil {
		log.Fatalln("Error getting contents")
	}

	return &t
}
