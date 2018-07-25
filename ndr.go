package stationcrawler

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/schnaidar/radiowatch"
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
	log := logrus.WithField("crawler", name)
	if err := readJson(url, &body); err != nil {
		return nil, err
	}

	trackInfos := strings.Split(body.SongNow, " - ")
	log.Debug(trackInfos)
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
