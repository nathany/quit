package mortgage

import (
	"testing"

	"github.com/nathany/quit/assert"
	"github.com/nathany/quit/float"
)

func TestEffectiveRate(t *testing.T) {
	result := float.Round(effectiveRate(0.0485), 5)
	expected := 0.04909
	assert.AssertInDelta(t, expected, result)

	result = effectiveRate(0.0234)
	expected = 0.02353689
	assert.AssertInDelta(t, expected, result)
}

func TestMonthlyRate(t *testing.T) {
	result := monthlyRate(effectiveRate(0.0234))
	expected := 0.001940561161
	assert.AssertInDelta(t, expected, result)
}

func TestMortgagePayment(t *testing.T) {
	assert.AssertInDelta(t, 2_703.90, MortgagePayment(700_000, 2.34, 30*12, Monthly))
	assert.AssertInDelta(t, 1_664.03, MortgagePayment(200_000, 5.85, 15*12, Monthly))

	// NOTE: Semimonthly payment matches the Royal Bank calculator,
	// but not the Government of Canada calculator ($831.39, that's strange).
	assert.AssertInDelta(t, 832.02, MortgagePayment(200_000, 5.85, 15*12, Semimonthly))
}
