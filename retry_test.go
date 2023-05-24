package retry

import (
  "errors"
  "testing"
  "time"
)

func alwaysFail() error {
  return errors.New("always fails")
}

func alwaysPass() error {
  return nil
}

func TestRetryFail(t *testing.T) {
  c := Config{
    MaxRetries: 3,
    Interval:   1 * time.Second,
  }

  err := Retry(alwaysFail, c)
  if err == nil {
    t.Error("expected an error, got nil")
  }
}

func TestRetrySuccess(t *testing.T) {
  c := Config{
    MaxRetries: 3,
    Interval:   1 * time.Second,
  }

  err := Retry(alwaysPass, c)
  if err != nil {
    t.Errorf("expected nil, got an error: %v", err)
  }
}

func TestRetryExponential(t *testing.T) {
  c := Config{
    MaxRetries:  3,
    Interval:    1 * time.Second,
    Exponential: true,
  }

  start := time.Now()

  err := Retry(alwaysFail, c)

  elapsed := time.Since(start)
  expectedElapsed := 7 * time.Second // 1s + 2s + 4s

  if err == nil || elapsed < expectedElapsed {
    t.Errorf("expected elapsed time to be at least %v, got %v", expectedElapsed, elapsed)
  }
}

func TestRetryMaxElapsedTime(t *testing.T) {
  c := Config{
    MaxRetries:     5,
    Interval:       2 * time.Second,
    MaxElapsedTime: 5 * time.Second,
  }

  start := time.Now()
  elapsedTime := time.Duration(0)

  err := Retry(func() error {
    elapsedTime = time.Since(start)

    // A function that always fails
    return errors.New("always fails")
  }, c)

  if elapsedTime > c.MaxElapsedTime {
    t.Errorf("expected elapsed time to be less than %v, got %v", c.MaxElapsedTime, elapsedTime)
  }

  if err == nil {
    t.Error("expected an error, got nil")
  }
}
