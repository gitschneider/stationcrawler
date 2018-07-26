package stationcrawler

import (
	"time"

	"github.com/gitschneider/radiowatch"
)

type Ndr1Crawler struct {
	crawler
}

func (n *Ndr1Crawler) Crawl() (*radiowatch.TrackInfo, error) {
	n.setNextCrawlTime(time.Now().Add(90 * time.Second))
	return crawlNdrStation("https://www.ndr.de/public/radioplaylists/ndr1niedersachsen.json", n.name)
}

func (n *Ndr1Crawler) Name() string {
	return n.name
}

func NewNdr1() *Ndr1Crawler {
	return &Ndr1Crawler{newCrawler("Ndr1 Niedersachsen")}
}
