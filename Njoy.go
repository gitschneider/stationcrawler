package stationcrawler

import (
	"time"

	"github.com/gitschneider/radiowatch"
)

type NjoyCrawler struct {
	crawler
}

func (n *NjoyCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	n.setNextCrawlTime(time.Now().Add(90 * time.Second))
	return crawlNdrStation("https://www.n-joy.de/public/radioplaylists/njoy.json", n.name)
}

func (n *NjoyCrawler) Name() string {
	return n.name
}

func NewNjoy() *NjoyCrawler {
	return &NjoyCrawler{newCrawler("N-Joy")}
}
