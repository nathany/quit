# When Can I Quit?

A library of Canadian 🇨🇦 personal finance calculations written in Go.

## Planned packages (subject to change)

This is work in progress.

* [x] assert - Helpful functions for testing
* [x] float - Helpful functions for working with floating point numbers
* [x] mortgage - Canadian mortgage calculations
  * [x] `MortgagePayment` calculation
  * [ ] Amortization schedule
  * [ ] Lump sum payments
* [x] taxes - Income tax calculations
  * [x] Very basic `IncomeTax` and `GrossIncomeForNet` calculations.
  * [ ] Tax brackets
  * [ ] Provinces and Territories other than Alberta
  * [ ] RRSP deduction calculations for contributions
  * Maybe CPP/EI calculations
  * Maybe self-employment vs. employment income
  * Maybe capital gains (in taxable accounts)
  * Maybe eligible and ineligible dividends (in taxable accounts)
* [ ] rrsp
  * [ ] RRSP withholding tax for withdrawals
  * [ ] RRSP contribution room calculations -- how many years to "max" it out?
* [ ] fi - Financial Independence calculations
  * [ ] The 4% rule
  * [ ] Coast FI calculator
  * [ ] Survival threshold from Die With Zero
* [ ] compound -- compound interest and drawdown calculators
  * [ ] Compound interest calculator for accumulation phase
  * [ ] Drawdown calculator
  * Maybe drawdown with CPP/OAS at different ages

### Maybe

There are some other calculations that may be interesting.

* Foreign withholding tax for dividends from U.S. and International equity
* T-Rex Scores (by Larry Bates) to compare fees on various investments, such as Management Expense Ratios (MER)
* Other suggestions?

## Disclaimer

The documentation is intended to describe how to use the functions in this library.

It is not financial advice.

## Contributing

Please open an issue to propose a change or addition before opening a pull request.

Please do NOT include personal information, financial or otherwise, in the code or examples. All numbers should be fictional.

## Resources

### Online Calculators

* [Mortgage Calculator](https://itools-ioutils.fcac-acfc.gc.ca/MC-CH/MCCalc-CHCalc-eng.aspx) by the Government of Canada
* [Income Tax Calculator](https://www.wealthsimple.com/en-ca/tool/tax-calculator) by Wealthsimple
* [Compound Interest Calculator](https://www.getsmarteraboutmoney.ca/calculators/compound-interest-calculator/) by Get Smarter About Money (Ontario Securities Commission)

### Blogs

* [A Guide to Canadian Mortgage Calculations (with code)](https://www.mikesukmanowsky.com/blog/a-guide-to-canadian-mortgage-calculations) by Mike Sukmanowsky
