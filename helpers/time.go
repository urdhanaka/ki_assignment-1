package helpers

import (
	"time"
)

func CalculateAlgorithmTime(start time.Time, end time.Time) uint64 {
	var timeElapsed uint64 = uint64(end.Sub(start).Nanoseconds())
	return timeElapsed
}
