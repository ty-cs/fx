package calculator

import "math"

// Rule72Result holds the results of a Rule of 72 calculation.
type Rule72Result struct {
	// One of Rate or Years will be the input; the other is computed.
	Rate       float64 // annual rate in percent
	Years      float64 // years to double
	RateMode   bool    // true = user gave rate, false = user gave years
	ExactYears float64 // exact years via ln(2)/ln(1+r) for comparison
	ExactRate  float64 // exact rate via ln(2)/years for comparison
}

// Rule72FromRate estimates how many years it takes to double at the given annual rate.
func Rule72FromRate(annualRate float64) Rule72Result {
	years72 := 72.0 / annualRate
	// Exact doubling time: ln(2) / ln(1 + r/100)
	exactYears := math.Log(2) / math.Log(1+annualRate/100)
	return Rule72Result{
		Rate:       annualRate,
		Years:      years72,
		RateMode:   true,
		ExactYears: exactYears,
	}
}

// Rule72FromYears estimates the rate needed to double in the given number of years.
func Rule72FromYears(years float64) Rule72Result {
	rate72 := 72.0 / years
	// Exact required rate: (2^(1/n) - 1) * 100
	exactRate := (math.Pow(2, 1/years) - 1) * 100
	return Rule72Result{
		Years:     years,
		Rate:      rate72,
		RateMode:  false,
		ExactRate: exactRate,
	}
}
