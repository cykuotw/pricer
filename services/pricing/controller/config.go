package controller

import (
	"pricing-app/config"
	"pricing-app/services/types"
)

func (c *Contoller) GetConfig(ticker string) (config.StockConfig, error) {
	cfg, ok := c.config[ticker]
	if !ok {
		return config.StockConfig{}, types.ErrInvalidTicker
	}

	return cfg, nil
}

func (c *Contoller) SetConfig(ticker string, cfg config.StockConfig) error {
	_, ok := c.config[ticker]
	if !ok {
		return types.ErrInvalidTicker
	}

	c.config[ticker] = cfg

	return nil
}
