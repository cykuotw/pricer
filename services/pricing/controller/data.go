package controller

import (
	"pricing-app/services/types"

	"github.com/shopspring/decimal"
)

// GetHistoryData retrieves the historical price data for a specific stock ticker.
//
// Parameters:
// - ticker: The stock ticker for which historical data is requested.
//
// Returns:
// - A slice of `decimal.Decimal` containing the historical prices.
// - An error if the ticker does not exist.
//
// Example usage:
//
//	prices, err := controller.GetHistoryData("AAPL")
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    log.Println("Historical Prices:", prices)
//	}
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
