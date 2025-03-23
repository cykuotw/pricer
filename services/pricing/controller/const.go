package controller

import "time"

var dt = 1.0 / (365 * 24 * 60) // 1 min as fraction of a year

// open time 9:30 AM
var MARKET_OPEN_TIME = time.Date(0, 1, 1, 9, 30, 0, 0, time.Local)

// close time 4:00 PM
var MARKET_CLOSE_TIME = time.Date(0, 1, 1, 16, 0, 0, 0, time.Local)
