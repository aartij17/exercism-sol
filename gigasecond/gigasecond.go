/*
Package gigasecond provides a simple utility which converts the given
timestamp into a timestamp in the future which has advanced by 10^9 seconds

Syntax of the time accepted is as follows:
	"2006-01-02"
	"2006-01-02T15:04:05"
*/
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond accepts a time object and adds 10^9 seconds to it and returns a new time object.
func AddGigasecond(t time.Time) time.Time {
	futureT := t.Add(time.Duration(int(math.Pow(10, 9))) * time.Second)
	return futureT
}
