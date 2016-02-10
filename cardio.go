package cardio

import "time"

type Tags map[string]string

type Values map[string]interface{}

type Beat struct {
	Name      string
	Timestamp time.Time
	Tags      Tags
	Values    Values
}

func NewBeat(name string) Beat {
	return NewBeatWithTS(name, time.Now())
}

func NewBeatWithTS(name string, when time.Time) Beat {
	return Beat{
		Name:      name,
		Timestamp: when,
		Tags:      make(Tags),
		Values:    make(Values),
	}
}

type Backend interface {
	Emit(Beat) error
}
