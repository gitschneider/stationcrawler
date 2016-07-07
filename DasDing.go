package stationcrawler

import (
	"time"
	"net/http"
	"github.com/schnaidar/radiowatch"
	"github.com/Jeffail/gabs"
	"errors"
	"fmt"
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
		// is there even a song being played?
		songTitle, ok := song.Search("title").Data().(string)
		if !ok {
			return nil, fmt.Errorf("No song is played at the moment!")
		}
		if songTitle == "" {
			return nil, fmt.Errorf("No song is played at the moment!")
		}

		scheduledAt, ok := song.Search("scheduledAtTs").Data().(string)
		if !ok {
			return nil, fmt.Errorf("Error while converting scheduledAt timestamp. Expected string, got %s",
				song.Search("scheduledAtTs").Data())
		}
		start, err := parseUnixString(scheduledAt)
		if err != nil {
			return nil, err
		}

		duration, ok := song.Search("duration").Data().(string)
		if !ok {
			return nil, fmt.Errorf("Error while converting duration. Expected tring, got %s",
				song.Search("duration").Data())
		}

		length, err := time.ParseDuration(duration + "s")
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
				Title: songTitle,
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