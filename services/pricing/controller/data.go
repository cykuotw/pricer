package controller

import (
	"pricing-app/services/types"

	"github.com/shopspring/decimal"
)

func (c *Contoller) GetHistoryData(ticker string) ([]decimal.Decimal, error) {
	exist := c.CheckTickerExist(ticker)
	if !exist {
		return nil, types.ErrInvalidTicker
	}

	tail := c.historyBufferTail[ticker]
	prices := make([]decimal.Decimal, tail)
	copy(prices, c.historyBuffer[ticker][:tail])

	return prices, nil
}
