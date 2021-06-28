package xtime

import (
	"errors"
	"time"
)

var idxOutOfRange = errors.New("goutils/xtime: GetDuration idx is out of range")

func IsIdxOutOfRangeErr(err error) bool {
	return errors.Is(err, idxOutOfRange)
}

func SingleStopwatch() *stopwatch {
	return newStopwatch(1)
}

func MultiStopwatch(cap int) *stopwatch {
	return newStopwatch(cap)
}

func newStopwatch(cap int) *stopwatch {
	return &stopwatch{
		st: time.Now(),
		et: make([]time.Duration, 0, cap),
	}
}

type stopwatch struct {
	st time.Time
	et []time.Duration
}

func (s *stopwatch) Stop() Duration {
	d := time.Since(s.st)
	s.et = append(s.et, d)
	return Duration(d)
}

func (s *stopwatch) GetDuration(idx int) (Duration, error) {
	if idx < 0 || idx >= len(s.et) {
		return 0, idxOutOfRange
	}

	return Duration(s.et[idx]), nil
}
