package controller

import "time"

// dt represents 1 minute as a fraction of a year.
// It is used in financial calculations for time-based simulations.
var dt = 1.0 / (365 * 24 * 60)

// MARKET_OPEN_TIME represents the market's opening time (9:30 AM).
// It is used to determine when the market starts operating each day.
var MARKET_OPEN_TIME = time.Date(0, 1, 1, 9, 30, 0, 0, time.Local)

// MARKET_CLOSE_TIME represents the market's closing time (4:00 PM).
// It is used to determine when the market stops operating each day.
var MARKET_CLOSE_TIME = time.Date(0, 1, 1, 16, 0, 0, 0, time.Local)
