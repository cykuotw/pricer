package controller

import (
	"pricing-app/services/types"
	"time"

	"github.com/shopspring/decimal"
)

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
