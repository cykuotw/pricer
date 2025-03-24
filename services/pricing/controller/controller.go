package controller

import (
	"pricing-app/config"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

// Contoller represents the core logic for managing stock price simulations.
type Contoller struct {
	config            config.MarketConfig          // Market configuration for all tickers.
	historyBuffer     map[string][]decimal.Decimal // Historical price buffer for each ticker.
	historyBufferTail map[string]int               // Tracks the tail index of the price buffer for each ticker.
}

// NewController creates a new instance of the Contoller.
// It initializes the historical price buffers and preloads simulated prices
// up to the current time if the market is open.
//
// Parameters:
// - cfg: The market configuration containing stock tickers and their initial settings.
//
// Returns:
// - A pointer to the initialized Contoller instance.
//
// Example usage:
//
//	marketConfig := config.LoadMarketConfig("data/initParam.json")
//	controller := NewController(marketConfig)
func NewController(cfg config.MarketConfig) *Contoller {
	openIntervalInMinute := 1 + int(MARKET_CLOSE_TIME.Sub(MARKET_OPEN_TIME).Minutes())
	udpateStartTime()
	now := time.Now()

	// Initialize buffers for historical prices and their tails.
	buffer := make(map[string][]decimal.Decimal)
	tails := make(map[string]int)
	for ticker, config := range cfg {
		buffer[ticker] = make([]decimal.Decimal, openIntervalInMinute)
		buffer[ticker][0] = decimal.NewFromUint64(config.Open)
		tails[ticker] = 1
	}

	// Generate prices if the engine starts within market hours.
	currTail := getCurrentTail(now)
	if currTail > openIntervalInMinute {
		currTail = openIntervalInMinute - 1
	}

	var wg sync.WaitGroup
	var m sync.Mutex

	// Preload prices for each ticker using concurrency.
	for ticker, conf := range cfg {
		wg.Add(1)

		go func(t string, c config.StockConfig) {
			for i := 1; i <= currTail; i++ {
				m.Lock()
				prevPrice := buffer[t][i-1]
				buffer[t][i] = simulateNextPrice(c, prevPrice)
				tails[t] = i + 1
				m.Unlock()
			}

		}(ticker, conf)
	}

	return &Contoller{
		config:            cfg,
		historyBuffer:     buffer,
		historyBufferTail: tails,
	}
}
