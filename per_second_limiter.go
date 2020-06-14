package ratelimiter

import (
	"context"
	"sync"
	"time"
)

//
func NewRateLimiterPerSecond(ctx context.Context, limit int) *SimplePerSecondRateLimiter {
	rl := &SimplePerSecondRateLimiter{rateLimit: limit, requests: make(map[string]int)}

	// reset limits
	timer := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				timer.Stop()
				rl.Stop()
				return
			case <-timer.C:
				rl.Reset()
			}
		}
	}()

	return rl
}

// SimplePerSecondRateLimiter implements mutex-based rate limiter.
type SimplePerSecondRateLimiter struct {
	ctx       context.Context
	rateLimit int
	requests  map[string]int
	rMux      sync.Mutex
}

func (rl *SimplePerSecondRateLimiter) AllowRequest(ip string) bool {
	rl.rMux.Lock()
	defer rl.rMux.Unlock()
	count, ok := rl.requests[ip]
	if !ok {
		rl.requests[ip] = 1
		ok = true
	}

	if ok && count < rl.rateLimit {
		rl.requests[ip] = count + 1
		return true
	}
	return false
}

func (rl *SimplePerSecondRateLimiter) Reset() {
	rl.rMux.Lock()
	defer rl.rMux.Unlock()
	rl.requests = nil
	rl.requests = make(map[string]int)
}

func (rl *SimplePerSecondRateLimiter) Stop() {
	rl.rMux.Lock()
	defer rl.rMux.Unlock()
	rl.requests = nil
}
