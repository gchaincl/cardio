package cardio

import "testing"

type testBackend struct {
	t     *testing.T
	check func(*testing.T, Beat)
}

func newTestBackend(t *testing.T, check func(*testing.T, Beat)) testBackend {
	return testBackend{t, check}
}

func (p testBackend) Emit(beat Beat) error {
	p.check(p.t, beat)
	return nil
}
