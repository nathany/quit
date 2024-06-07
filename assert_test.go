package main

import "testing"

func AssertInDelta(t *testing.T, expected, actual float64) {
	if !inDelta(expected, actual) {
		t.Errorf("Expected %f, but got %f", expected, actual)
	}
}
