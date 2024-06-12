package float

import "testing"

func AssertInDelta(t *testing.T, expected, actual float64) {
	if !InDelta(expected, actual, Epsilon) {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}
