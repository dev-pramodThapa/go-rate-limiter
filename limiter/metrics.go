package limiter

import (
	"sync"
)

// RequestMetrics holds metrics for each request.
type RequestMetrics struct {
	RequestsCount     int `json:"request_count"`
	ThrottledRequests int `json:"throttled_requests"`
}

// Metrics manages metrics for various requests.
type Metrics struct {
	mu      sync.RWMutex
	metrics map[string]*RequestMetrics
}

// NewMetrics initializes a new Metrics instance with an empty map.
func NewMetrics() *Metrics {
	return &Metrics{
		metrics: make(map[string]*RequestMetrics),
	}
}

// IncreaseRequestsCount increments the total request count for a given endpoint.
func (m *Metrics) IncreaseRequestsCount(endpoint string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	metrics, exists := m.metrics[endpoint]
	if !exists {
		metrics = &RequestMetrics{}
		m.metrics[endpoint] = metrics
	}
	metrics.RequestsCount++
}

// IncreaseThrottledRequest increments the count of throttled requests for a given endpoint.
func (m *Metrics) IncreaseThrottledRequest(endpoint string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	metrics, exists := m.metrics[endpoint]
	if !exists {
		metrics = &RequestMetrics{}
		m.metrics[endpoint] = metrics
	}
	metrics.ThrottledRequests++
}

// GetMetrics returns a copy of the metrics map to prevent external modifications.
func (m *Metrics) GetMetrics() map[string]*RequestMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Create a new map to hold the copied metrics
	copyMetrics := make(map[string]*RequestMetrics, len(m.metrics))
	for k, v := range m.metrics {
		copyMetrics[k] = &RequestMetrics{
			RequestsCount:     v.RequestsCount,
			ThrottledRequests: v.ThrottledRequests,
		}
	}

	return copyMetrics
}
