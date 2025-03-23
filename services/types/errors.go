package types

import "errors"

var (
	ErrEmptyRequestBody = errors.New("missing request body")

	ErrInvalidTicker = errors.New("Invalid ticker")
)
