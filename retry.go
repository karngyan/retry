package retry

import (
	"errors"
	"time"
)

// RetryableFunc is a function type for functions that can be retried
type RetryableFunc func() error

// Config holds the configuration for retries
type Config struct {
	MaxRetries     int
	Interval       time.Duration
	Exponential    bool
	MaxElapsedTime time.Duration
}

// Retry retries a function based on the provided configuration
func Retry(f RetryableFunc, c Config) error {
	var err error
	elapsedTime := time.Duration(0)

	for i := 0; i < c.MaxRetries; i++ {
		err = f()
		if err == nil {
			return nil
		}

		// Sleep for the interval time
		time.Sleep(c.Interval)

		elapsedTime += c.Interval
		if c.MaxElapsedTime > 0 && elapsedTime > c.MaxElapsedTime {
			return errors.New("maximum elapsed time for retries has been exceeded")
		}

		// Increase the interval exponentially if configured to do so
		if c.Exponential {
			c.Interval *= 2
		}
	}

	return err
}
