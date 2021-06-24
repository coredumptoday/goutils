package timex

import "time"

type timeDev time.Duration

func (d timeDev) Seconds() float64 {
	return time.Duration(d).Seconds()
}

func (d timeDev) Milliseconds() int64 {
	return time.Duration(d).Milliseconds()
}

func (d timeDev) ExactMilliseconds() float64 {
	return float64(d) / 1e6
}

func (d timeDev) Microseconds() int64 {
	return time.Duration(d).Microseconds()
}

func (d timeDev) ExactMicroseconds() float64 {
	return float64(d) / 1e3
}

func (d timeDev) Nanoseconds() int64 {
	return time.Duration(d).Nanoseconds()
}
