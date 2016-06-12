package stationcrawler


func NewHr3() *hrCrawler {
	return &hrCrawler{newCrawler("hr3"), "http://playlist.hr-online.de/playlist/hr3box.php"}
}