package limiter

import (
	"go-rate-limiter/utils"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens     int
	maxTokens  int
	lastRefill time.Time
	refillRate time.Duration
	mu         sync.Mutex
}

// NewTokenBucket creates a new TokenBucket with specified maxTokens and refillRate.
func NewTokenBucket(maxTokens int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// IsRequestAllowed checks if a request is allowed and updates the token bucket.
func (tb *TokenBucket) IsRequestAllowed() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	
	tb.refillTokens()

	// Check if there are tokens available and decrement if so
	if tb.tokens > 0 {
		tb.tokens--

		return true
	}

	return false
}

// refillTokens updates the number of tokens in the bucket based on the elapsed time.
func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	// Calculate the number of tokens to add based on elapsed time
	newTokens := int(elapsed / tb.refillRate)
	tb.tokens = utils.Min(tb.tokens+newTokens, tb.maxTokens)
	tb.lastRefill = now
}

// ResetBucket resets the token bucket for a given ID and bucket type with updated settings.
func (rl *RateLimiter) ResetBucket(id, bucketType string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	config := rl.getRateLimitConfig(bucketType, id)

	switch bucketType {
	case "users":
		rl.userBuckets[id] = NewTokenBucket(config.MaxTokens, config.RefillRate)
	case "admin":
		rl.adminBuckets[id] = NewTokenBucket(config.MaxTokens, config.RefillRate)
	}
}

// GetBucket returns the token bucket for the given ID and bucket type. If it does not exist, a new one is created.
func (rl *RateLimiter) GetBucket(id, bucketType string, maxTokens int, refillRate time.Duration) *TokenBucket {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	var bucket *TokenBucket
	switch bucketType {
	case "users":
		if _, exists := rl.userBuckets[id]; !exists {
			rl.userBuckets[id] = NewTokenBucket(maxTokens, refillRate)
		}
		bucket = rl.userBuckets[id]
	case "admin":
		if _, exists := rl.adminBuckets[id]; !exists {
			rl.adminBuckets[id] = NewTokenBucket(maxTokens, refillRate)
		}
		bucket = rl.adminBuckets[id]
	}

	return bucket
}
