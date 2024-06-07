package main

import (
	"math"
)

const epsilon = 1e-9

func inDelta(x, y float64) bool {
	return math.Abs(x-y) < epsilon
}

func round(val float64, precision int) float64 {
	return math.Round(val*(math.Pow10(precision))) / math.Pow10(precision)
}

func roundUp(val float64, precision int) float64 {
	return math.Ceil(val*(math.Pow10(precision))) / math.Pow10(precision)
}
