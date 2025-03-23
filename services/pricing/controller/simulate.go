package controller

import (
	"math"
	"math/rand/v2"
	"pricing-app/config"

	"github.com/shopspring/decimal"
)

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
