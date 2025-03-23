package controller

import (
	"time"
)

func (c *Contoller) UpdatePriceToLatestMin(ticker string, now time.Time) {
	currTail := int(now.Sub(MARKET_OPEN_TIME).Minutes())
	conf := c.config[ticker]
	tail := c.historyBufferTail[ticker]

	for i := tail; i <= currTail; i++ {
		prevPrice := c.historyBuffer[ticker][i-1]
		c.historyBuffer[ticker][i] = simulateNextPrice(conf, prevPrice)
		c.historyBufferTail[ticker] = i + 1
	}
}
