package stationcrawler

import (
	"time"

	"github.com/gitschneider/radiowatch"
)

type iLoveRadioSongInfo struct {
	Channel_1 struct {
		Artist    string `json:"artist"`
		ChannelID string `json:"channel_id"`
		Color     string `json:"color"`
		Cover     string `json:"cover"`
		Fontcolor string `json:"fontcolor"`
		Title     string `json:"title"`
	} `json:"channel-1"`
}

type IloveRadioCrawler struct {
	crawler
}

func (d *IloveRadioCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	var body iLoveRadioSongInfo
	if err := readJson("https://iloveradio.de/typo3conf/ext/ep_channel/Scripts/playlist.php", &body); err != nil {
		return nil, err
	}

	d.setNextCrawlTime(time.Now().Add(90 * time.Second))
	return &radiowatch.TrackInfo{
		Artist:    body.Channel_1.Artist,
		Title:     body.Channel_1.Title,
		Station:   d.Name(),
		CrawlTime: time.Now(),
	}, nil
}

func (d *IloveRadioCrawler) Name() string {
	return d.name
}

func NewIloveRadio() *IloveRadioCrawler {
	return &IloveRadioCrawler{newCrawler("I Love Radio")}
}
