package controller

import (
	"time"
)

func udpateStartTime() {
	now := time.Now()

	MARKET_OPEN_TIME = time.Date(now.Year(), now.Month(), now.Day(), MARKET_OPEN_TIME.Hour(), MARKET_OPEN_TIME.Minute(), MARKET_OPEN_TIME.Second(), MARKET_OPEN_TIME.Nanosecond(), MARKET_OPEN_TIME.Location())
	MARKET_CLOSE_TIME = time.Date(now.Year(), now.Month(), now.Day(), MARKET_CLOSE_TIME.Hour(), MARKET_CLOSE_TIME.Minute(), MARKET_CLOSE_TIME.Second(), MARKET_CLOSE_TIME.Nanosecond(), MARKET_CLOSE_TIME.Location())
}

func (c *Contoller) CheckMarketOpen(now time.Time) bool {
	return now.After(MARKET_OPEN_TIME) && now.Before(MARKET_CLOSE_TIME)
}

func (c *Contoller) GetMarketTime() (time.Time, time.Time) {
	return MARKET_OPEN_TIME, MARKET_CLOSE_TIME
}

func getCurrentTail(now time.Time) int {
	// +0.01 hear is to avoid floating point error:
	// 		i.e.
	// 		234.99985375111666 is considered 235,
	// 		but the casting will ignore the trailing 0.9998...
	// 		by adding 0.01 is a hacky way to solve this issue

	return int(now.Sub(MARKET_OPEN_TIME).Minutes() + 0.1)
}
