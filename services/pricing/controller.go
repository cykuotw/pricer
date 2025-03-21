package pricing

import "pricing-app/config"

type Contoller struct {
	config config.MarketConfig
}

func NewController(cfg config.MarketConfig) *Contoller {
	return &Contoller{
		config: cfg,
	}
}
