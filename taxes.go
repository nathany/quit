package main

// 2024 tax rates for Alberta
// via https://www.taxtips.ca/taxrates/ab.htm
// and https://www.taxtips.ca/taxrates/canada.htm
const (
	provincialBasicPersonalAmount = 21_885
	provincialTaxRate             = 0.1 // 10% for the first $148,269

	federalBasicPersonalAmount = 15_705 // for income of $173,205 or less
	federalTaxRate             = 0.15   // 15% for the first $55,867

	maxGrossIncome = 55_867.00
)

// IncomeTax calculates federal and provincial income tax for gross amounts less than $55,867.
func IncomeTax(income float64) float64 {
	return round(federalTax(income)+provincialTax(income), 2)
}

func federalTax(income float64) float64 {
	if income <= federalBasicPersonalAmount {
		return 0
	}
	return (income - federalBasicPersonalAmount) * federalTaxRate
}

// TODO: implement tax brackets if needed
func provincialTax(income float64) float64 {
	if income <= provincialBasicPersonalAmount {
		return 0
	}
	return (income - provincialBasicPersonalAmount) * provincialTaxRate
}

func GrossIncomeForNet(desiredNet float64) float64 {
	// brute force approach scans for the correct net amount
	for gross := desiredNet; gross <= maxGrossIncome; gross += 0.01 {
		tax := IncomeTax(gross)
		net := round(gross-tax, 2)
		if inDelta(net, desiredNet) {
			return round(gross, 2)
		}
	}
	return 0
}
