package stationcrawler

func NewYouFm() *hrCrawler {
	return &hrCrawler{newCrawler("YOU FM"), "http://www.you-fm.de/playlist/youfmbox.php"}
}
