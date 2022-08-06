package retry_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/swathinsankaran/retry"
)

func TestRetry(t *testing.T) {
	var count int
	err := retry.Do(func() (bool, error) {
		count++
		return true, nil
	}, 10, 1)
	if err != nil {
		t.Fatalf("Failed: expected %v, but got %v", nil, err)
	}
}

func TestRetryTimeExceeded(t *testing.T) {
	err := retry.Do(func() (bool, error) {
		time.Sleep(2 * time.Millisecond)
		return true, nil
	}, 10, 1)
	t.Logf("Got an error: %v", err)
	if !errors.Is(err, retry.ErrExceededRetryTimeout) {
		t.Fatalf("Failed: expected %v, but got %v", retry.ErrExceededRetryTimeout, err)
	}
}

func TestRetryMaxAttemptExceeded(t *testing.T) {
	err := retry.Do(func() (bool, error) {
		return false, nil
	}, 10, 1)
	t.Logf("Got an error: %v", err)
	if !errors.Is(err, retry.ErrMaxRetryAttemptExceeded) {
		t.Fatalf("Failed: expected %v, but got %v", retry.ErrMaxRetryAttemptExceeded, err)
	}
}

func TestRetryOnFuncErr(t *testing.T) {
	errExpected := fmt.Errorf("error string")
	err := retry.Do(func() (bool, error) {
		return false, errExpected
	}, 10, 1)
	t.Logf("Got an error: %v", err)
	if !errors.Is(err, errExpected) {
		t.Fatalf("Failed: expected %v, but got %v", errExpected, err)
	}
}
