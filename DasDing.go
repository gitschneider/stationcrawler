package stationcrawler

import (
	"time"
	"net/http"
	"github.com/schnaidar/radiowatch"
	"github.com/Jeffail/gabs"
	"errors"
)

type dasDingSongInfo struct {
	Artist        string `json:"artist"`
	ArtistDetails interface{} `json:"artistDetails"`
	CoverURL      string `json:"coverUrl"`
	CurrentLikes  string `json:"currentLikes"`
	DetailPageURL string `json:"detailPageUrl"`
	Duration      string `json:"duration"`
	HookFile      string `json:"hookFile"`
	ID            string `json:"id"`
	LikeURL       string `json:"likeUrl"`
	PlayedAtTs    string `json:"playedAtTs"`
	ScheduledAtTs string `json:"scheduledAtTs"`
	SocialMessage string `json:"socialMessage"`
	Title         string `json:"title"`
}

type DasDingCrawler struct {
	crawler
}

func (d *DasDingCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	resp, err := http.Get("http://www.dasding.de/ext/playlist/gen2_pl.json")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	jsonParsed, err := gabs.ParseJSONBuffer(resp.Body)
	if err != nil {
		return nil, err
	}
	songs, err := jsonParsed.Children()
	if err != nil {
		return nil, err
	}

	var song *gabs.Container
	for _, song = range songs {
		start, err := parseUnixString(song.Search("scheduledAtTs").Data().(string))
		if err != nil {
			return nil, err
		}

		length, err := time.ParseDuration(song.Search("duration").Data().(string) + "s")
		if err != nil {
			return nil, err
		}

		end := start.Add(length)
		if isNow(start, end) {
			d.setNextCrawlTime(end)
			return &radiowatch.TrackInfo{
				Artist:song.Search("artist").Data().(string),
				CrawlTime: time.Now(),
				Station: d.name,
				Title: song.Search("title").Data().(string),
			}, nil
		}
	}

	return nil, errors.New("No song is played at the moment.")
}

func (d *DasDingCrawler)Name() string {
	return d.name
}

func NewDasDing() *DasDingCrawler {
	return &DasDingCrawler{newCrawler("Das Ding")}
}