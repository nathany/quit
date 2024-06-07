package main

import "math"

type PaymentFrequency int

const (
	Monthly     PaymentFrequency = 12
	Semimonthly PaymentFrequency = 24
	Biweekly    PaymentFrequency = 26
	Weekly      PaymentFrequency = 52
)

// Helpful resource for understanding Canadian mortgage calculations:
// https://www.mikesukmanowsky.com/blog/a-guide-to-canadian-mortgage-calculations

// Calculate effective interest rate or  from the nominal rate.
// Or Annual Percentage Yield (APY) from the Annual Percentage Rate (APR).
// nominalRate should be a decimal value, e.g. 0.0485 for 4.85%
// source: https://www.aprtoapy.com/
func effectiveRate(nominalRate float64) float64 {
	// Canadian fixed-rate mortgages compound semi-annually (2 periods/year)
	// So does my variable-rate mortgage, though that isn't always the case.
	const periods = 2
	return math.Pow(1+(nominalRate/periods), periods) - 1
}

func monthlyRate(effectiveRate float64) float64 {
	return periodicRate(effectiveRate, 12)
}

// calculate periodic interest rate (e.g semi-monthly is 24 periods)
func periodicRate(effectiveRate float64, periods int) float64 {
	return math.Pow(1+effectiveRate, 1.0/float64(periods)) - 1
}

func MortgagePayment(principalAmount, rateAsPercent, amortizationMonths float64, frequency PaymentFrequency) float64 {
	monthlyRate := monthlyRate(effectiveRate(rateAsPercent / 100))
	monthlyPayment := monthlyRate * principalAmount / (1 - math.Pow(1+monthlyRate, -amortizationMonths))
	return round(monthlyPayment*12/float64(frequency), 2)
}
