package controller

import (
	"time"
)

// udpateStartTime updates the `MARKET_OPEN_TIME` and `MARKET_CLOSE_TIME` variables
// to reflect the current date while keeping the same time of day.
// This ensures that the market times are correctly set for the current day.
func udpateStartTime() {
	now := time.Now()

	MARKET_OPEN_TIME = time.Date(now.Year(), now.Month(), now.Day(), MARKET_OPEN_TIME.Hour(), MARKET_OPEN_TIME.Minute(), MARKET_OPEN_TIME.Second(), MARKET_OPEN_TIME.Nanosecond(), MARKET_OPEN_TIME.Location())
	MARKET_CLOSE_TIME = time.Date(now.Year(), now.Month(), now.Day(), MARKET_CLOSE_TIME.Hour(), MARKET_CLOSE_TIME.Minute(), MARKET_CLOSE_TIME.Second(), MARKET_CLOSE_TIME.Nanosecond(), MARKET_CLOSE_TIME.Location())
}

// CheckMarketOpen determines whether the market is currently open.
//
// Parameters:
// - now: The current time to check against the market's open and close times.
//
// Returns:
// - A boolean indicating whether the market is open.
func (c *Contoller) CheckMarketOpen(now time.Time) bool {
	return now.After(MARKET_OPEN_TIME) && now.Before(MARKET_CLOSE_TIME)
}

// GetMarketTime retrieves the market's opening and closing times.
//
// Returns:
// - Two `time.Time` values representing the market's opening and closing times.
func (c *Contoller) GetMarketTime() (time.Time, time.Time) {
	return MARKET_OPEN_TIME, MARKET_CLOSE_TIME
}

// getCurrentTail calculates the current tail index for the historical price buffer
// based on the time elapsed since the market opened.
//
// Parameters:
// - now: The current time to calculate the tail index.
//
// Returns:
// - An integer representing the current tail index.
//
// Note:
// - A small value (0.1) is added to avoid floating-point errors during calculations.
func getCurrentTail(now time.Time) int {
	// +0.1 is added to avoid floating-point errors, ensuring accurate rounding.
	// 		i.e.
	// 		234.99985375111666 is considered 235,
	// 		but the casting will ignore the trailing 0.9998...
	// 		by adding 0.01 is a hacky way to solve this issue

	return int(now.Sub(MARKET_OPEN_TIME).Minutes() + 0.1)
}
