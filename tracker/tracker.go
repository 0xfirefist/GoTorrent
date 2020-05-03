package tracker

import (
	"github.com/jackpal/bencode-go"
	"log"
	"net/http"
	"time"
)

type TrackerResponse struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
	Failure  string `bencode:"failure reason, omitempty"`
}

func GetResponse(url string) *TrackerResponse {
	c := &http.Client{Timeout: 15 * time.Second}
	resp, err := c.Get(url)
	if err != nil {
		log.Fatalf("get response failed")
	}
	defer resp.Body.Close()

	tresp := &TrackerResponse{}
	err = bencode.Unmarshal(resp.Body, &tresp)
	if err != nil {
		log.Fatalf("error getting tracker response - %s", err)
	}
	if tresp.Failure != "" {
		log.Fatalf("query failed")
	}

	return tresp
}
