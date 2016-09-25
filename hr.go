package stationcrawler

import (
	"time"
	"github.com/schnaidar/radiowatch"
	"errors"
	"html"
)

type hrSongInfo struct {
	Songs  []struct {
		Dauer     string `json:"dauer"`
		Idsong    int    `json:"idsong"`
		Interpret string `json:"interpret"`
		Realzeit  int    `json:"realzeit"`
		Titel     string `json:"titel"`
		Zeit      string `json:"zeit"`
	} `json:"songs"`
	Update int `json:"update"`
}

type hrCrawler struct {
	crawler
	url string
}

func (d *hrCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	var body hrSongInfo
	err := readJson(d.url, &body)
	if err != nil {
		return nil, err
	}

	for _, song := range body.Songs {
		start := time.Unix(int64(song.Realzeit), 0)
		dauer, err := time.ParseDuration(song.Dauer + "s")
		if err != nil {
			return nil, err
		}
		end := start.Add(dauer)
		if isNow(start, end) {
			d.setNextCrawlTime(end.Add(10 * time.Second))
			return &radiowatch.TrackInfo{
				Artist: html.UnescapeString(body.Songs[3].Interpret),
				CrawlTime: time.Now(),
				Station: d.name,
				Title: html.UnescapeString(body.Songs[3].Titel),
			}, nil
		}
	}
	return nil, errors.New("No song is played at the moment.")
}

func (d *hrCrawler)Name() string {
	return d.name
}