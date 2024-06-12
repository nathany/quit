/*
This is an income tax calculator for Alberta, Canada with data for 2024.

It doesn't currently handle tax brackets, as it's currently intended
for calculating taxes during retirement and assumes <= $55,867 gross income/year.
*/
package taxes

import "github.com/nathany/quit/float"

// 2024 Tax rates for Alberta
// via https://www.taxtips.ca/taxrates/ab.htm
// and https://www.taxtips.ca/taxrates/canada.htm
const (
	// TODO: this could be expanded to support other provinces
	provincialBasicPersonalAmount = 21_885
	provincialTaxRate             = 0.1 // 10% for the first $148,269

	federalBasicPersonalAmount = 15_705 // for income of $173,205 or less
	federalTaxRate             = 0.15   // 15% for the first $55,867

	maxGrossIncome = 55_867.00
)

// IncomeTax calculates federal and provincial income tax for gross amounts less than $55,867.
func IncomeTax(income float64) float64 {
	return float.Round(federalTax(income)+provincialTax(income), 2)
}

func federalTax(income float64) float64 {
	// TODO: implement tax brackets
	if income <= federalBasicPersonalAmount {
		return 0
	}
	return (income - federalBasicPersonalAmount) * federalTaxRate
}

func provincialTax(income float64) float64 {
	// TODO: implement tax brackets
	if income <= provincialBasicPersonalAmount {
		return 0
	}
	return (income - provincialBasicPersonalAmount) * provincialTaxRate
}

// GrossIncomeForNet can help determine an RRSP drawdown amount given a desired net amount for cost of living.
func GrossIncomeForNet(desiredNet float64) float64 {
	// brute force approach scans for the correct net amount
	for gross := desiredNet; gross <= maxGrossIncome; gross += 0.01 {
		tax := IncomeTax(gross)
		net := float.Round(gross-tax, 2)
		if float.InDelta(net, desiredNet, float.Epsilon) {
			return float.Round(gross, 2)
		}
	}
	return 0
}
