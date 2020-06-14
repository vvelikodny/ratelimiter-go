package ratelimiter

type rateLimiter interface {
	RequestAllowed(ip string) bool
	ResetLimits()
	Stop()
}
