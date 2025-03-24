package controller

import (
	"pricing-app/services/types"
	"time"
)

// CheckPriceUpdated checks if the price for a given ticker has been updated
// based on the current time.
//
// Parameters:
// - ticker: The stock ticker to check.
// - now: The current time to compare against the last update.
//
// Returns:
// - A boolean indicating whether the price has been updated.
// - An error if the ticker does not exist.
//
// Example usage:
//
//	updated, err := controller.CheckPriceUpdated("AAPL", time.Now())
//	if err != nil {
//	    log.Println("Error:", err)
//	}
//	if updated {
//	    log.Println("Price has been updated.")
//	}
func (c *Contoller) CheckPriceUpdated(ticker string, now time.Time) (bool, error) {
	if exist := c.CheckTickerExist(ticker); !exist {
		return false, types.ErrInvalidTicker
	}

	tail := c.historyBufferTail[ticker]

	currTail := getCurrentTail(now)

	return currTail < tail, nil
}
