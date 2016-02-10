package logger

import (
	"fmt"
	"time"

	"github.com/gchaincl/cardio"
)

type Beat cardio.Beat

func (b Beat) String() string {
	return fmt.Sprintf(
		"%s (%s) [%v] <%v>\n",
		b.Name,
		b.Timestamp.Format(time.Stamp),
		b.Tags,
		b.Values,
	)
}

type Backend struct{}

func New() Backend {
	return Backend{}
}

func (Backend) Emit(beat cardio.Beat) error {
	_, err := fmt.Printf("%s", Beat(beat))
	return err
}
