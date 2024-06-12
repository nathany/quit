/*
Assert package contains helpful functions for testing.
*/
package assert

import (
	"testing"

	"github.com/nathany/quit/float"
)

func AssertInDelta(t *testing.T, expected, actual float64) {
	if !float.InDelta(expected, actual, float.Epsilon) {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}
