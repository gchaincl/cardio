package cardio

import (
	"runtime"
	"time"
)

type Pulse struct {
	name    string
	backend Backend
	tick    time.Duration
	cancel  chan struct{}
}

func NewPulse(name string, backend Backend) Pulse {
	pulse := Pulse{
		name:    name,
		backend: backend,
		tick:    1 * time.Second,
		cancel:  make(chan struct{}),
	}
	go pulse.loop()

	return pulse
}

func (p Pulse) Tick(tick time.Duration) Pulse {
	p.Cancel()

	p.tick = tick
	go p.loop()

	return p
}

func (p Pulse) loop() {
	for {
		select {
		case <-time.After(p.tick):
			p.backend.Emit(
				newPulseBeat(p.name),
			)
		case <-p.cancel:
			return
		}
	}
}

func (p Pulse) Cancel() {
	p.cancel <- struct{}{}
}

func newPulseBeat(name string) Beat {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	ngr := runtime.NumGoroutine()

	beat := NewBeat(name)
	beat.Values = map[string]interface{}{
		"num_goroutines": ngr,
		"alloc":          mem.Alloc,
		"total_alloc":    mem.TotalAlloc,
	}

	return beat
}
