package cardio

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPulseEmitTheRightBeatFormat(t *testing.T) {
	done := make(chan struct{})

	NewPulse("test", newTestBackend(t,
		func(t *testing.T, beat Beat) {
			assert.Equal(t, "test", beat.Name)
			assert.Equal(t, runtime.NumGoroutine(), beat.Values["num_goroutines"])
			assert.NotZero(t, beat.Values["alloc"])
			assert.NotZero(t, beat.Values["total_alloc"])

			done <- struct{}{}
		}),
	).Tick(1 * time.Millisecond)

	<-done
}

func TestPulseIsCanceled(t *testing.T) {
	failed := false

	NewPulse("test", newTestBackend(t,
		func(t *testing.T, beat Beat) {
			if !failed {
				assert.Fail(t, "Canceled Pulse should not emit any beat")
			}
			failed = true
		}),
	).Tick(1 * time.Millisecond).Cancel()

	time.Sleep(2 * time.Millisecond)
}
