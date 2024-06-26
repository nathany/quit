/*
Float package contains helpful functions for working with floating point numbers.

# A Note on Floating Point Numbers

Floating point numbers are a poor choice for financial calculations due to rounding errors.
However, the calculations in this module are intended for rough estimates rather than precise accounting.
*/
package float

import (
	"math"
)

const (
	Epsilon = 1e-9
	Penny   = 0.01
)

func InDelta(x, y, delta float64) bool {
	return math.Abs(x-y) < delta
}

func Round(val float64, precision int) float64 {
	return math.Round(val*(math.Pow10(precision))) / math.Pow10(precision)
}

func RoundUp(val float64, precision int) float64 {
	return math.Ceil(val*(math.Pow10(precision))) / math.Pow10(precision)
}
