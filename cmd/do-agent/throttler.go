package main

import "time"

type constThrottler struct {
	wait time.Duration
}

func (c *constThrottler) WaitDuration() time.Duration {
	_ = "STUB: not implemented"

	// Name is the name of this limiter
	return *new(time.Duration)
}

func (c *constThrottler) Name() string { _ = "STUB: not implemented"; return "" }
