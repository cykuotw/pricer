package controller

import "time"

func udpateStartTime() {
	now := time.Now()

	MARKET_OPEN_TIME = time.Date(now.Year(), now.Month(), now.Day(), MARKET_OPEN_TIME.Hour(), MARKET_OPEN_TIME.Minute(), MARKET_OPEN_TIME.Second(), MARKET_OPEN_TIME.Nanosecond(), MARKET_OPEN_TIME.Location())
	MARKET_CLOSE_TIME = time.Date(now.Year(), now.Month(), now.Day(), MARKET_CLOSE_TIME.Hour(), MARKET_CLOSE_TIME.Minute(), MARKET_CLOSE_TIME.Second(), MARKET_CLOSE_TIME.Nanosecond(), MARKET_CLOSE_TIME.Location())
}

func checkMarketOpen(now time.Time) bool {
	return now.After(MARKET_OPEN_TIME) && now.Before(MARKET_CLOSE_TIME)
}

func (c *Contoller) CheckMarketOpen(now time.Time) bool {
	return now.After(MARKET_OPEN_TIME) && now.Before(MARKET_CLOSE_TIME)
}

func (c *Contoller) GetMarketTime() (time.Time, time.Time) {
	return MARKET_OPEN_TIME, MARKET_CLOSE_TIME
}
