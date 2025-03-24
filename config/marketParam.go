package config

import (
	"encoding/json"
	"os"

	"github.com/shopspring/decimal"
)

// StockConfig represents the configuration for a single stock.
type StockConfig struct {
	Open       uint64          `json:"open"`       // Price when the market is open
	Drift      decimal.Decimal `json:"drift"`      // Drift value for the stock's price simulation.
	Volatility decimal.Decimal `json:"volatility"` // Volatility value for the stock's price simulation.
}

// MarketConfig represents the configuration for the entire market.
type MarketConfig map[string]StockConfig

// LoadMarketConfig loads the market configuration from a JSON file.
// It takes the file path as input and returns a MarketConfig object or an error.
//
// Example usage:
//
//	config, err := LoadMarketConfig("data/initParam.json")
//	if err != nil {
//	    log.Fatal(err)
//	}
func LoadMarketConfig(filePath string) (MarketConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config MarketConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return config, nil
}
