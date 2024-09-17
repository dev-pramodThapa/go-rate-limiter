package limiter

import (
	"sync"
	"time"
)

type TokenBucketConfig struct {
	MaxTokens  int           `json:"max_tokens"`
	RefillRate time.Duration `json:"refill_rate"`
}

type RateLimitConfig struct {
	mu          sync.RWMutex
	RateLimiter *RateLimiter
	Metrics     *Metrics
	UserConfig  map[string]TokenBucketConfig
	AdminConfig map[string]TokenBucketConfig
}

// AppConfig is a global variable holding the rate limit configuration.
var AppConfig *RateLimitConfig

// GetConfig returns the global rate limit configuration.
func GetConfig() *RateLimitConfig {
	return AppConfig
}

// InitConfig initializes and returns a new RateLimitConfig with default settings.
func InitConfig() *RateLimitConfig {
	return &RateLimitConfig{
		RateLimiter: NewRateLimiter(),
		Metrics:     NewMetrics(),
		UserConfig: map[string]TokenBucketConfig{
			"default": {MaxTokens: 5, RefillRate: time.Minute / 5},
		},
		AdminConfig: map[string]TokenBucketConfig{
			"default": {MaxTokens: 2, RefillRate: time.Minute / 2},
		},
	}
}

// UpdateConfig updates the rate limit configuration for a given endpoint type and ID.
func (c *RateLimitConfig) UpdateConfig(endpointType, id string, maxTokens int, refillRate time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	config := TokenBucketConfig{MaxTokens: maxTokens, RefillRate: refillRate}

	switch endpointType {
	case "users":
		c.UserConfig[id] = config
	case "admin":
		c.AdminConfig[id] = config
	}
}

// GetConfig retrieves the rate limit configuration for a given endpoint type and ID.
func (c *RateLimitConfig) GetConfig(endpointType, id string) TokenBucketConfig {
	c.mu.RLock()
	defer c.mu.RUnlock()

	switch endpointType {
	case "users":
		if config, exists := c.UserConfig[id]; exists {
			return config
		}
		return c.UserConfig["default"]
	case "admin":
		if config, exists := c.AdminConfig[id]; exists {
			return config
		}
		return c.AdminConfig["default"]
	default:
		return TokenBucketConfig{}
	}
}
