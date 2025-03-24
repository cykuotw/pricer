package controller

import (
	"pricing-app/config"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type Contoller struct {
	config            config.MarketConfig
	historyBuffer     map[string][]decimal.Decimal
	historyBufferTail map[string]int
}

func NewController(cfg config.MarketConfig) *Contoller {
	openIntervalInMinute := 1 + int(MARKET_CLOSE_TIME.Sub(MARKET_OPEN_TIME).Minutes())
	udpateStartTime()
	now := time.Now()

	// initialize buffer
	buffer := make(map[string][]decimal.Decimal)
	tails := make(map[string]int)
	for ticker, config := range cfg {
		buffer[ticker] = make([]decimal.Decimal, openIntervalInMinute)
		buffer[ticker][0] = decimal.NewFromUint64(config.Open)
		tails[ticker] = 1
	}

	// generate prices if engine start within opening hour
	currTail := getCurrentTail(now)
	if currTail > openIntervalInMinute {
		currTail = openIntervalInMinute - 1
	}

	var wg sync.WaitGroup
	var m sync.Mutex

	for ticker, conf := range cfg {
		wg.Add(1)

		// optimized with multi-threading
		// just in case there are too many ticker to update
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
