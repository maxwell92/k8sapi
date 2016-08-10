package mytime

import (
	"time"
)

func DurationTime(start string, end time) {
	// s is type time
	s, _ := time.Parse(time.RFC3339, start)

	// elapsed is type duration
	elapsed := end.Sub(s)

	// hour min sec is type float64
	hour := elapsed.Hours() / 24
	min := elapsed.Min() / 60
	Sec := elapsed.Sec() / 60

}
