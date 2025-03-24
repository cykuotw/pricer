package controller

import (
	"pricing-app/config"
	"pricing-app/services/types"
)

// GetConfig retrieves the configuration for a specific stock ticker.
//
// Parameters:
// - ticker: The stock ticker for which the configuration is requested.
//
// Returns:
// - A `config.StockConfig` object containing the stock's configuration.
// - An error if the ticker does not exist.
//
// Example usage:
//
//	cfg, err := controller.GetConfig("AAPL")
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    log.Println("Config:", cfg)
//	}
func (c *Contoller) GetConfig(ticker string) (config.StockConfig, error) {
	cfg, ok := c.config[ticker]
	if !ok {
		return config.StockConfig{}, types.ErrInvalidTicker
	}

	return cfg, nil
}

// SetConfig updates the configuration for a specific stock ticker.
//
// Parameters:
// - ticker: The stock ticker for which the configuration is being updated.
// - cfg: A `config.StockConfig` object containing the new configuration values.
//
// Returns:
// - An error if the ticker does not exist.
//
// Example usage:
//
//	err := controller.SetConfig("AAPL", newConfig)
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    log.Println("Config updated successfully.")
//	}
func (c *Contoller) SetConfig(ticker string, cfg config.StockConfig) error {
	_, ok := c.config[ticker]
	if !ok {
		return types.ErrInvalidTicker
	}

	c.config[ticker] = cfg

	return nil
}
