package limiter

import (
	"sync"
)

type RateLimiter struct {
	mu           sync.Mutex
	userBuckets  map[string]*TokenBucket
	adminBuckets map[string]*TokenBucket
}

// NewRateLimiter initializes a new RateLimiter with empty token buckets.
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		userBuckets:  make(map[string]*TokenBucket),
		adminBuckets: make(map[string]*TokenBucket),
	}
}

// getRateLimitConfig retrieves the rate limit configuration based on the bucket type and ID.
func (rl *RateLimiter) getRateLimitConfig(bucketType, id string) TokenBucketConfig {
	switch bucketType {
		
	case "user":
		return AppConfig.GetConfig("user", id)
	case "admin":
		return AppConfig.GetConfig("admin", id)
	default:
		return TokenBucketConfig{}
	}
}