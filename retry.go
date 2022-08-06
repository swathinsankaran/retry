package retry

import (
	"context"
	"errors"
	"time"
)

var (
	ErrExceededRetryTimeout    = errors.New("exceeded retry timeout")
	ErrMaxRetryAttemptExceeded = errors.New("exceeded maximum retry attempts")
)

// Func represents the function to retry.
type Func func() (bool, error)

// Do retries the function passed for `maxRetryAttempt` attempts or
// cancels the execution when the time exceeds `retryTimeout` milliseconds.
func Do(f Func, maxRetryAttempt int, retryTimeout int) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(retryTimeout*int(time.Millisecond)),
	)
	defer cancel()
	chErr := make(chan error)
	go func() {
		attempt := 1
		for {
			success, err := f()
			if err != nil || success {
				chErr <- err
				return
			}
			if attempt == maxRetryAttempt {
				cancel()
				return
			}
			attempt++
		}
	}()
	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			return ErrExceededRetryTimeout
		case context.Canceled:
			return ErrMaxRetryAttemptExceeded
		}
	case err := <-chErr:
		return err
	}
	return nil
}
