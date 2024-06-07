package main

import "testing"

func TestEffectiveRate(t *testing.T) {
	result := round(effectiveRate(0.0485), 5)
	expected := 0.04909
	AssertInDelta(t, expected, result)

	result = effectiveRate(0.0234)
	expected = 0.02353689
	AssertInDelta(t, expected, result)
}

func TestMonthlyRate(t *testing.T) {
	result := monthlyRate(effectiveRate(0.0234))
	expected := 0.001940561161
	AssertInDelta(t, expected, result)
}

func TestMortgagePayment(t *testing.T) {
	AssertInDelta(t, 2_703.90, MortgagePayment(700_000, 2.34, 30*12, Monthly))
	AssertInDelta(t, 1_664.03, MortgagePayment(200_000, 5.85, 15*12, Monthly))
	AssertInDelta(t, 832.01, MortgagePayment(200_000, 5.85, 15*12, Semimonthly))
}
