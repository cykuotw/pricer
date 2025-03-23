package controller

import (
	"pricing-app/services/types"
	"time"
)

func (c *Contoller) CheckPriceUpdated(ticker string, now time.Time) (bool, error) {
	if exist := c.CheckTickerExist(ticker); !exist {
		return false, types.ErrInvalidTicker
	}

	tail := c.historyBufferTail[ticker]
	currTail := int(now.Sub(MARKET_OPEN_TIME).Minutes())

	return currTail == tail, nil
}
