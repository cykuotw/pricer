package config

import (
	"encoding/json"
	"os"

	"github.com/shopspring/decimal"
)

type StockConfig struct {
	Open       uint64          `json:"open"`
	Drift      decimal.Decimal `json:"drift"`
	Volatility decimal.Decimal `json:"volatility"`
}

type MarketConfig map[string]StockConfig

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
