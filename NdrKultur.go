package stationcrawler

import (
	"time"

	"github.com/gitschneider/radiowatch"
)

type NdrKulturCrawler struct {
	crawler
}

func (n *NdrKulturCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	n.setNextCrawlTime(time.Now().Add(90 * time.Second))
	return crawlNdrStation("https://www.ndr.de/public/radioplaylists/ndrkultur.json", n.name)
}

func (n *NdrKulturCrawler) Name() string {
	return n.name
}

func NewNdrKultur() *NdrKulturCrawler {
	return &NdrKulturCrawler{newCrawler("NDR Kultur")}
}
