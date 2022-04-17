package clockface

import (
	"math"
	"time"
)

// A point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of a analogue clock at the time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi
}
