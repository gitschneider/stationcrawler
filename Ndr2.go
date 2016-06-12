package stationcrawler

import (
	"github.com/schnaidar/radiowatch"
	"time"
)

type Ndr2Crawler struct {
	crawler
}

func (n *Ndr2Crawler) Crawl() (*radiowatch.TrackInfo, error) {
	n.setNextCrawlTime(time.Now().Add(90 * time.Second))
	return crawlNdrStation("https://www.ndr.de/public/radio_playlists/ndr2.json", n.name)
}

func (n *Ndr2Crawler) Name() string {
	return n.name
}

func NewNdr2() *Ndr2Crawler {
	return &Ndr2Crawler{newCrawler("Ndr2")}
}
