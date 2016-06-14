package stationcrawler

import (
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"github.com/schnaidar/radiowatch"
	"time"
	"fmt"
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

	songStart, err := time.Parse("2006-01-02 15:04:05", songInfo.Onair.Start)
	if err != nil {
		return nil, err
	}

	songDuration, err := time.ParseDuration(songInfo.Onair.Duration + "ms")
	if err != nil {
		return nil, err
	}
	c.setNextCrawlTime(songStart.Add(songDuration))
	fmt.Println(songStart)
	fmt.Println(c.NextCrawlTime())
	return &radiowatch.TrackInfo{
		Artist: songInfo.Onair.Artist,
		Title: songInfo.Onair.Title,
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