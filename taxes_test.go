package main

import (
	"testing"
)

func TestIncomeTax(t *testing.T) {
	result := IncomeTax(44_000)

	if result != 6_455.75 {
		t.Errorf("Expected $6,455.75, but got %f", result)
	}
}

func TestGrossIncomeForNet(t *testing.T) {
	result := GrossIncomeForNet(37_544.25)
	if result != 44_000 {
		t.Errorf("Expected $44,000, but got %f", result)
	}

	result = GrossIncomeForNet(30_000.00)
	if result != 33_941 {
		t.Errorf("Expected $33,941, but got %f", result)
	}
}
