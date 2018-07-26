package stationcrawler

func NewYouFm() *hrCrawler {
	return &hrCrawler{newCrawler("YOU FM"), "http://playlist.hr-online.de/playlist/youfmbox.php"}
}
