package ratelimiter

import (
	"context"
	"testing"
	"time"
)

const (
	RequestsPerSecondLimit = 100
)

func TestMiddlewarePositive(t *testing.T) {
	ip := "testIP"

	ctx, cancel := context.WithCancel(context.Background())

	middleware := NewMockMiddleware(NewRateLimiterPerSecond(ctx, RequestsPerSecondLimit))

	for i := 0; i < RequestsPerSecondLimit; i++ {
		if err := middleware.Proceed(ip); err != nil {
			t.FailNow()
		}
		time.Sleep(time.Microsecond * 100)
	}

	cancel()
}

func TestMiddlewareNegative(t *testing.T) {
	ip := "testIP"

	ctx, cancel := context.WithCancel(context.Background())
	middleware := NewMockMiddleware(NewRateLimiterPerSecond(ctx, RequestsPerSecondLimit))

	var stateErr error
	for i := 0; i < RequestsPerSecondLimit; i++ {
		stateErr = middleware.Proceed(ip)
		time.Sleep(time.Millisecond * 15)
	}

	if stateErr != nil {
		t.FailNow()
	}

	cancel()
}
