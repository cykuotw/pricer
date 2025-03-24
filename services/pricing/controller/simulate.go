package controller

import (
	"math"
	"math/rand/v2"
	"pricing-app/config"

	"github.com/shopspring/decimal"
)

// simulateNextPrice calculates the next simulated price for a stock based on its configuration
// and the previous price using a Geometric Brownian Motion model.
//
// Parameters:
// - config: A `config.StockConfig` object containing the stock's drift and volatility values.
// - prevPrice: The previous price of the stock as a `decimal.Decimal`.
//
// Returns:
// - The next simulated price as a `decimal.Decimal`, rounded to 2 decimal places.
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

// randNorm generates a random number from a standard normal distribution
// using the Box-Muller Transform.
//
// Returns:
// - A random float64 value from a standard normal distribution.
func randNorm() float64 {
	// Reference: https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform

	u1 := rand.Float64()
	u2 := rand.Float64()

	return math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2.0*math.Pi*u2)
}
