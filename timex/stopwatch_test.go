package timex

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSingleStopwatch(t *testing.T) {
	ss := SingleStopwatch()
	time.Sleep(time.Millisecond * 2344)
	fmt.Println(ss.Stop().Seconds())
	fmt.Println(ss.Stop().Milliseconds())
	fmt.Println(ss.Stop().ExactMilliseconds())
	fmt.Println(ss.Stop().Microseconds())
	fmt.Println(ss.Stop().ExactMicroseconds())
}

func TestMultiStopwatch(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ss := MultiStopwatch(3)
	time.Sleep(time.Millisecond * 1111)
	fmt.Println(ss.Stop().ExactMilliseconds())
	time.Sleep(time.Millisecond * 1111)
	fmt.Println(ss.Stop().ExactMilliseconds())
	time.Sleep(time.Millisecond * 1111)
	fmt.Println(ss.Stop().ExactMilliseconds())
	time.Sleep(time.Millisecond * 1111)
	fmt.Println(ss.Stop().ExactMilliseconds())
}
