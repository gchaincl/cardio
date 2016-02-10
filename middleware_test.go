package cardio

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMiddlewareTrackStatusAndUrl(t *testing.T) {
	hf := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(418)
	}

	mw := NewMiddleware("test", newTestBackend(t,
		func(t *testing.T, beat Beat) {
			assert.Equal(t, "418", beat.Tags["status"])
			assert.Equal(t, "/testing_cardio", beat.Tags["path"])
		}),
	).Handler

	ts := httptest.NewServer(mw(http.HandlerFunc(hf)))
	defer ts.Close()

	http.Get(ts.URL + "/testing_cardio")
}

func TestMiddlewareTrackRequestTime(t *testing.T) {
	var begin time.Time

	hf := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
	}

	mw := NewMiddleware("test", newTestBackend(t,
		func(t *testing.T, beat Beat) {
			req_time := beat.Values["request_time"].(time.Duration)

			assert.True(t, req_time > 10*time.Millisecond)
			assert.True(t, req_time < time.Now().Sub(begin))
		}),
	).Handler

	ts := httptest.NewServer(mw(http.HandlerFunc(hf)))
	defer ts.Close()

	begin = time.Now()
	http.Get(ts.URL)
}
