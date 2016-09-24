package stationcrawler

import (
	"strings"
	"errors"
	"time"
	"github.com/schnaidar/radiowatch"
	"html"
)

type ndrSongInfo struct {
	Action       string `json:"action"`
	NextVisitIn  string `json:"nextVisitIn"`
	SongNext     string `json:"song_next"`
	SongNow      string `json:"song_now"`
	SongNowCover string `json:"song_now_cover"`
	SongPrevious string `json:"song_previous"`
	TimeStamp    string `json:"timeStamp"`
}

func crawlNdrStation(url string, name string) (*radiowatch.TrackInfo, error) {
	var body ndrSongInfo
	readJson(url, &body)

	trackInfos := strings.Split(body.SongNow, " - ")
	if len(trackInfos) != 2 {
		return nil, errors.New("Did not get info about current track!")
	}

	return &radiowatch.TrackInfo{
		Artist:    html.UnescapeString(trackInfos[0]),
		Title:     html.UnescapeString(trackInfos[1]),
		Station:   name,
		CrawlTime: time.Now(),
	}, nil
}
