package controller

import (
	"pricing-app/services/types"
	"time"

	"github.com/shopspring/decimal"
)

// UpdatePriceToLatestMin updates the price of a specific stock ticker to the latest minute.
// It generates simulated prices for all missing minutes up to the current time.
//
// Parameters:
// - ticker: The stock ticker to update.
// - now: The current time to calculate the latest minute.
//
// Returns:
// - An error if the ticker does not exist or if the market is closed.
//
// Example usage:
//
//	err := controller.UpdatePriceToLatestMin("AAPL", time.Now())
//	if err != nil {
//	    log.Println("Error updating price:", err)
//	}
func (c *Contoller) UpdatePriceToLatestMin(ticker string, now time.Time) error {
	if exist := c.CheckTickerExist(ticker); !exist {
		return types.ErrInvalidTicker
	}

	if updated, _ := c.CheckPriceUpdated(ticker, now); updated {
		return nil
	}

	currTail := getCurrentTail(now)
	conf := c.config[ticker]
	tail := c.historyBufferTail[ticker]

	if currTail > len(c.historyBuffer[ticker]) {
		// the market is closed
		return nil
	}

	for i := tail; i <= currTail; i++ {
		prevPrice := c.historyBuffer[ticker][i-1]
		c.historyBuffer[ticker][i] = simulateNextPrice(conf, prevPrice)
		c.historyBufferTail[ticker] = i + 1
	}

	return nil
}

// UpdatePrice updates the price of a specific stock ticker to the current time.
// If the price is already updated, it returns the latest price without recalculating.
//
// Parameters:
// - ticker: The stock ticker to update.
// - now: The current time to calculate the latest price.
//
// Returns:
// - The updated price as a `decimal.Decimal`.
// - An error if the ticker does not exist.
//
// Example usage:
//
//	price, err := controller.UpdatePrice("AAPL", time.Now())
//	if err != nil {
//	    log.Println("Error updating price:", err)
//	} else {
//	    log.Println("Updated price:", price)
//	}
func (c *Contoller) UpdatePrice(ticker string, now time.Time) (decimal.Decimal, error) {
	if exist := c.CheckTickerExist(ticker); !exist {
		return decimal.Decimal{}, types.ErrInvalidTicker
	}

	if updated, _ := c.CheckPriceUpdated(ticker, now); updated {
		return c.historyBuffer[ticker][c.historyBufferTail[ticker]-1], nil
	}

	conf := c.config[ticker]
	tail := c.historyBufferTail[ticker]

	prevPrice := c.historyBuffer[ticker][tail-1]
	c.historyBuffer[ticker][tail] = simulateNextPrice(conf, prevPrice)
	c.historyBufferTail[ticker] = tail + 1

	return c.historyBuffer[ticker][c.historyBufferTail[ticker]-1], nil
}
