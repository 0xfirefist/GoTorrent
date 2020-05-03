package torrent

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/jackpal/bencode-go"
)

func Open(path string) *torrent {

	// file read
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("file not found - %s", path)
	}

	// unmarshal
	t := torrent{}
	err = bencode.Unmarshal(bytes.NewReader(fileContent), &t)
	if err != nil {
		log.Fatalln("error Unmarshalling - %s ", err)
	}

	return &t
}
