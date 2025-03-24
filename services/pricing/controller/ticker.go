package controller

import "sort"

// GetTickers retrieves a sorted list of all available stock tickers.
//
// Returns:
// - A slice of strings containing all stock tickers in alphabetical order.
//
// Example usage:
//
//	tickers := controller.GetTickers()
//	log.Println("Available tickers:", tickers)
func (c *Contoller) GetTickers() []string {
	tickers := make([]string, 0, len(c.config))
	for ticker := range c.config {
		tickers = append(tickers, ticker)
	}
	sort.Strings(tickers)

	return tickers
}

// CheckTickerExist checks if a specific stock ticker exists in the market configuration.
//
// Parameters:
// - ticker: The stock ticker to check.
//
// Returns:
// - A boolean indicating whether the ticker exists.
//
// Example usage:
//
//	exists := controller.CheckTickerExist("AAPL")
//	if exists {
//	    log.Println("Ticker exists.")
//	} else {
//	    log.Println("Ticker does not exist.")
//	}
func (c *Contoller) CheckTickerExist(ticker string) bool {
	_, ok := c.config[ticker]

	return ok
}
