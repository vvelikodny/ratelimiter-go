package ratelimiter

type rateLimiter interface {
	AllowRequest(ip string) bool
	Reset()
	Stop()
}
