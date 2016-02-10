package statsd

import "github.com/gchaincl/cardio"

type Backend struct {
}

func New(addr string) Backend {
	return Backend{}
}

func (b Backend) Emit(cardio.Beat) error {
	return nil
}
