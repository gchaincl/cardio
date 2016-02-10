package cardio

import (
	"net/http"
	"strconv"
	"time"
)

type httpWriter struct {
	http.ResponseWriter
	wroteHeader bool
	status      int
}

func (w *httpWriter) WriteHeader(status int) {
	if !w.wroteHeader {
		w.status = status
		w.wroteHeader = true
		w.ResponseWriter.WriteHeader(status)
	}
}

func (w httpWriter) Status() int {
	return w.status
}

type Middleware struct {
	name    string
	backend Backend
}

func NewMiddleware(name string, backend Backend) Middleware {
	return Middleware{
		name:    name,
		backend: backend,
	}
}

func (mw Middleware) Handler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		hw := &httpWriter{ResponseWriter: w}

		t1 := time.Now()
		h.ServeHTTP(hw, r)

		if hw.Status() == 0 {
			hw.WriteHeader(http.StatusOK)
		}
		t2 := time.Now()

		mw.backend.Emit(
			newHTTPBeat(mw.name, r.URL.Path, t2.Sub(t1), hw.Status()),
		)
	}

	return http.HandlerFunc(fn)
}

func newHTTPBeat(name string, path string, req_time time.Duration, status int) Beat {
	beat := NewBeat(name)
	beat.Tags["path"] = path
	beat.Tags["status"] = strconv.Itoa(status)
	beat.Values["request_time"] = req_time

	return beat
}
