package ratelimit

import (
	"time"
)

type RateLimit struct {
	Timeframe time.Duration // Timeframe as a duration (e.g., 10*time.Second)
	MaxCalls  int           // Maximum number of calls within the timeframe
}
