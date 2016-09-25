package stationcrawler

import (
	"github.com/schnaidar/radiowatch"
	"errors"
	"time"
	"html"
)

type AntenneSongInfo struct {
	Result struct {
			   Entry []struct {
				   Airtime  string `json:"airtime"`
				   Duration string `json:"duration"`
				   Song     struct {
								Entry []struct {
									Artist struct {
											   Entry []struct {
												   GUID string `json:"guid"`
												   Name string `json:"name"`
											   } `json:"entry"`
											   Found string `json:"found"`
										   } `json:"artist"`
									GUID   string `json:"guid"`
									Title  string `json:"title"`
								} `json:"entry"`
								Found string `json:"found"`
							} `json:"song"`
			   } `json:"entry"`
			   Found string `json:"found"`
		   } `json:"result"`
}

type AntenneCrawler struct {
	crawler
}

func (c *AntenneCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	var song AntenneSongInfo
	err := readJson("http://webradio.antenne.com/plugin/data.php?data=microflow&station=3&offset=1&count=1", &song)
	if err != nil {
		return nil, err
	}

	if song.Result.Found == "0" {
		return nil, errors.New("No Song is played at the moment!")
	}

	start, err := time.Parse("2006-01-02T15:04:05-07:00", song.Result.Entry[0].Airtime)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(song.Result.Entry[0].Duration + "s")
	if err != nil {
		return nil, err
	}

	c.setNextCrawlTime(start.Add(duration))
	return &radiowatch.TrackInfo{
		Title: html.UnescapeString(song.Result.Entry[0].Song.Entry[0].Title),
		Artist: html.UnescapeString(song.Result.Entry[0].Song.Entry[0].Artist.Entry[0].Name),
		Station: c.Name(),
		CrawlTime: time.Now(),
	}, nil
}

func (c	*AntenneCrawler) Name() string {
	return c.name
}

func NewAntenne() *AntenneCrawler{
	return &AntenneCrawler{newCrawler("Antenne Niedersachsen")}
}