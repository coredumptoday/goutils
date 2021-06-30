package time

import "time"

type Duration time.Duration

func (d Duration) Seconds() float64 {
	return time.Duration(d).Seconds()
}

func (d Duration) Milliseconds() int64 {
	return time.Duration(d).Milliseconds()
}

func (d Duration) ExactMilliseconds() float64 {
	return float64(d) / 1e6
}

func (d Duration) Microseconds() int64 {
	return time.Duration(d).Microseconds()
}

func (d Duration) ExactMicroseconds() float64 {
	return float64(d) / 1e3
}

func (d Duration) Nanoseconds() int64 {
	return time.Duration(d).Nanoseconds()
}
