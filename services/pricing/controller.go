package pricing

import (
	"math"
	"math/rand/v2"
	"pricing-app/config"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

var dt = 1.0 / (365 * 24 * 60) // 1 min as fraction of a year

// open time 9:30 AM
var MARKET_OPEN_TIME = time.Date(0, 1, 1, 9, 30, 0, 0, time.Local)

// close time 4:00 PM
var MARKET_CLOSE_TIME = time.Date(0, 1, 1, 16, 0, 0, 0, time.Local)

type Contoller struct {
	config            config.MarketConfig
	historyBuffer     map[string][]decimal.Decimal
	historyBufferTail map[string]int
}

func NewController(cfg config.MarketConfig) *Contoller {
	openIntervalInMinute := int(MARKET_CLOSE_TIME.Sub(MARKET_OPEN_TIME).Minutes())
	now := time.Now()
	// now = time.Date(0, 1, 1, now.Hour(), now.Minute(), now.Second(), 0, time.Local)
	now = time.Date(0, 1, 1, 12, 0, 0, 0, time.Local) // debug

	// initialize buffer
	buffer := make(map[string][]decimal.Decimal)
	tails := make(map[string]int)
	for ticker, config := range cfg {
		buffer[ticker] = make([]decimal.Decimal, openIntervalInMinute)
		buffer[ticker][0] = decimal.NewFromUint64(config.Open)
		tails[ticker] = 1
	}

	// generate prices if engine start within opening hour
	if now.After(MARKET_OPEN_TIME) && now.Before(MARKET_CLOSE_TIME) {
		var wg sync.WaitGroup
		var m sync.Mutex
		currTail := int(now.Sub(MARKET_OPEN_TIME).Minutes())

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
	}

	return &Contoller{
		config:            cfg,
		historyBuffer:     buffer,
		historyBufferTail: tails,
	}
}

func simulateNextPrice(config config.StockConfig, prevPrice decimal.Decimal) decimal.Decimal {
	prev, _ := prevPrice.Float64()
	mu, _ := config.Drift.Float64()
	sigma, _ := config.Volatility.Float64()
	z := randNorm()

	exponent := (mu-0.5*sigma*sigma)*dt + sigma*math.Sqrt(dt)*z
	nextPrice := prev * math.Exp(exponent)
	roundedNextPrice := decimal.NewFromFloat(nextPrice).Round(2)

	return roundedNextPrice
}

func randNorm() float64 {
	// Box Muller Transform
	// ref: https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform

	u1 := rand.Float64()
	u2 := rand.Float64()

	return math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2.0*math.Pi*u2)
}
