package stationcrawler

import (
	"fmt"
	"time"
	"strings"
	"errors"
	"html"
	"github.com/schnaidar/radiowatch"
)

type mdrJumpSongInfo struct {
	Info       string `json:"Info"`
	Resulttype string `json:"Resulttype"`
	Songs      struct {
		           Zero struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"0"`
		           One struct {
			                ArtistImageID struct {
				                              _root        string        `json:"@root"`
				                              ImageVariant []interface{} `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"1"`
		           Two struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"2"`
		           Three struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"3"`
		           Four struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"4"`
		           Five struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"5"`
		           Six struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"6"`
		           Seven struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"7"`
		           Eight struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"8"`
		           Nine struct {
			                ArtistImageID struct {
				                              _root        string `json:"@root"`
				                              ImageVariant []struct {
					                              _attributes struct {
						                                          Height   string `json:"height"`
						                                          MimeType string `json:"mimeType"`
						                                          Name     string `json:"name"`
						                                          URL      string `json:"url"`
						                                          Width    string `json:"width"`
					                                          } `json:"@attributes"`
				                              } `json:"imageVariant"`
			                              } `json:"artist_image_id"`
			                Audioasset struct {
				                              _root string        `json:"@root"`
				                              Asset []interface{} `json:"asset"`
			                              } `json:"audioasset"`
			                Author           string `json:"author"`
			                AvID             string `json:"av_id"`
			                AvNextID         string `json:"av_next_id"`
			                Duration         string `json:"duration"`
			                IDTitel          string `json:"id_titel"`
			                Interpret        string `json:"interpret"`
			                InterpretURL     string `json:"interpret_url"`
			                Komponist        string `json:"komponist"`
			                Kurzinfo         string `json:"kurzinfo"`
			                Label            string `json:"label"`
			                Metadatentext    string `json:"metadatentext"`
			                Starttime        string `json:"starttime"`
			                Status           string `json:"status"`
			                Subtitle         string `json:"subtitle"`
			                Title            string `json:"title"`
			                Tontraeger       string `json:"tontraeger"`
			                Transmissiontype string `json:"transmissiontype"`
		                } `json:"9"`
	           } `json:"Songs"`
}

type MdrJumpCrawler struct {
	crawler
}

func (c *MdrJumpCrawler) Crawl() (*radiowatch.TrackInfo, error) {
	var song mdrJumpSongInfo
	url := fmt.Sprintf("http://www.jumpradio.de/XML/titellisten/jump_onair.json?json&_=%v", time.Now().Unix())
	err := readJson(url, &song)
	if err != nil {
		return nil, err
	}

	if song.Resulttype != "OK" {
		return nil, NoSongNowError()
	}

	start, err := time.Parse("2006-01-02 15:04:05", song.Songs.Zero.Starttime)
	if err != nil {
		return nil, err
	}

	durationParts := strings.Split(song.Songs.Zero.Duration, ":")
	durationString := fmt.Sprintf("%vh%vm%vs", durationParts[0], durationParts[1], durationParts[2])
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return nil, errors.New("Error while parse song duration")
	}

	c.setNextCrawlTime(start.Add(duration))
	return &radiowatch.TrackInfo{
		Artist: html.UnescapeString(song.Songs.Zero.Interpret),
		Title: html.UnescapeString(song.Songs.Zero.Title),
		CrawlTime: time.Now(),
		Station: c.Name(),
	}, nil
}

func (c *MdrJumpCrawler) Name() string {
	return c.name
}

func NewMdrJump() *MdrJumpCrawler {
	return &MdrJumpCrawler{newCrawler("MDR JUMP")}
}
