package stationcrawler

import "time"

type crawler struct {
	name          string
	// The time at which this crawler should be run next time
	nextCrawlTime time.Time
}

func (c *crawler) NextCrawlTime() time.Time {
	return c.nextCrawlTime
}

func (c *crawler) setNextCrawlTime(t time.Time) {
	c.nextCrawlTime = t
}

func newCrawler(name string) crawler {
	return crawler{name, time.Now().Add(-1 * time.Minute)}
}