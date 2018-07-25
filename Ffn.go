package stationcrawler

import (
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"github.com/gitschneider/radiowatch"
	"time"
	"html"
)

// XML? Srsly?

type ffnSongInfo struct {
	Onair struct {
			  Artist   string `xml:"artist"`
			  Title    string `xml:"title"`
			  Start    string `xml:"start"`
			  Duration string `xml:"duration"`
		  } `xml:"onair"`
}

type FfnCrawler struct {
	crawler
}

func (c *FfnCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	resp, err := http.Get("http://www.ffn.de/fileadmin/content/ffn/ffn.xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var songInfo ffnSongInfo
	err = xml.Unmarshal(bodyBytes, &songInfo)
	if err != nil {
		return nil, err
	}

	// Falling back to periodical check as the duration property seems to be weird.
	// The value translates most of the time to about 5 minutes, so some songs are skipped when crawling.
	c.setNextCrawlTime(time.Now().Add(90 * time.Second))

	return &radiowatch.TrackInfo{
		Artist: html.UnescapeString(songInfo.Onair.Artist),
		Title: html.UnescapeString(songInfo.Onair.Title),
		CrawlTime: time.Now(),
		Station: c.Name(),
	}, nil
}

func NewFfn() *FfnCrawler {
	return &FfnCrawler{newCrawler("FFN")}
}

func (c *FfnCrawler) Name() string {
	return c.name
}