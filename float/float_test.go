package float

import "testing"

func TestInDeltaEpsilon(t *testing.T) {
	if !InDelta(1.0, 1.0000000001, Epsilon) {
		t.Errorf("Expected 1.0, 1.0000000001 to be InDelta, but got false")
	}
	if InDelta(1.0, 1.000000001, Epsilon) {
		t.Errorf("Expected 1.0, 1.000000001 to not be InDelta, but got true")
	}
}

func TestInDeltaPenny(t *testing.T) {
	if !InDelta(1.0, 1.005, Penny) {
		t.Errorf("Expected 1.0, 1.005 to be InDelta, but got false")
	}
	if InDelta(1.0, 1.01, Penny) {
		t.Errorf("Expected 1.0, 1.01 to not be InDelta, but got true")
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		input     float64
		precision int
		expected  float64
	}{
		{1.23456, 2, 1.23},
		{1.23556, 2, 1.24},
		{1.23456, 3, 1.235},
	}

	for _, test := range tests {
		result := Round(test.input, test.precision)
		if result != test.expected {
			t.Errorf("Expected %f to round to %f, but got %f", test.input, test.expected, result)
		}
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		input     float64
		precision int
		expected  float64
	}{
		{1.23456, 2, 1.24},
		{1.23556, 2, 1.24},
		{1.23414, 3, 1.235},
		{1.234, 3, 1.234},
	}

	for _, test := range tests {
		result := RoundUp(test.input, test.precision)
		if result != test.expected {
			t.Errorf("Expected %f to round to %f, but got %f", test.input, test.expected, result)
		}
	}
}
