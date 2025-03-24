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

	currTail := getCurrentTail(now)

	return currTail == tail-1, nil
}
