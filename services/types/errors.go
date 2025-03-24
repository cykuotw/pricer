package types

import "errors"

var (
	// ErrEmptyRequestBody is returned when the request body is missing or empty.
	ErrEmptyRequestBody = errors.New("missing request body")

	// ErrInvalidTicker is returned when an invalid or non-existent stock ticker is provided.
	ErrInvalidTicker = errors.New("Invalid ticker")
)
