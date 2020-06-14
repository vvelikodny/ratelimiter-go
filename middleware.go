package ratelimiter

import (
	"errors"
	"fmt"
)

var (
	rateLimitReached = errors.New("rate limit reached")
)

type Middleware interface {
	Proceed(ip string) error
}

type middlewareMock struct {
	rateLimiter rateLimiter
}

// NewMockMiddleware creates instance of middlewareMock to test rate limiter.
func NewMockMiddleware(limiter rateLimiter) Middleware {
	return &middlewareMock{
		rateLimiter: limiter,
	}
}

// Proceed checks if request from ip could be processed.
func (m *middlewareMock) Proceed(ip string) error {
	if !m.rateLimiter.RequestAllowed(ip) {
		return fmt.Errorf("proceed request from ip %s: %w", ip, rateLimitReached) // ex. HTTP 429
	}

	// call next() middleware
	return nil
}
