package ratelimiter

import (
	"errors"
	"fmt"
)

type Middleware interface {
	Proceed(ip string) error
}

var (
	rateLimitReached = errors.New("rate limit reached")
)

type middlewareMock struct {
	rateLimiter rateLimiter
}

func NewMockMiddleware(limiter rateLimiter) Middleware {
	return &middlewareMock{
		rateLimiter: limiter,
	}
}

func (m *middlewareMock) Proceed(ip string) error {
	if !m.rateLimiter.AllowRequest(ip) {
		return fmt.Errorf("proceed request from ip %s: %w", ip, rateLimitReached) // ex. HTTP 429
	}

	// call next() middleware
	return nil
}
