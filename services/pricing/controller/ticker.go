package controller

import "sort"

func (c *Contoller) GetTickers() []string {
	tickers := make([]string, 0, len(c.config))
	for ticker := range c.config {
		tickers = append(tickers, ticker)
	}
	sort.Strings(tickers)

	return tickers
}

func (c *Contoller) CheckTickerExist(ticker string) bool {
	_, ok := c.config[ticker]

	return ok
}
