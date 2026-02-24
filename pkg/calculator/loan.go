package calculator

import "math"

// LoanResult holds the results of a loan calculation.
type LoanResult struct {
	Principal      float64
	AnnualRate     float64
	Years          int
	MonthlyPayment float64
	TotalPayment   float64
	TotalInterest  float64
	Schedule       []AmortizationRow
}

// AmortizationRow represents one month of the amortization schedule.
type AmortizationRow struct {
	Month     int
	Payment   float64
	Principal float64
	Interest  float64
	Balance   float64
}

// CalculateLoan computes the monthly payment and full amortization schedule.
func CalculateLoan(principal, annualRate float64, years int) LoanResult {
	months := years * 12
	monthlyRate := annualRate / 100 / 12

	var monthlyPayment float64
	if monthlyRate == 0 {
		// 0% interest: evenly split
		monthlyPayment = principal / float64(months)
	} else {
		// Standard annuity formula: P * r(1+r)^n / ((1+r)^n - 1)
		factor := math.Pow(1+monthlyRate, float64(months))
		monthlyPayment = principal * monthlyRate * factor / (factor - 1)
	}

	var schedule []AmortizationRow
	balance := principal
	for m := 1; m <= months; m++ {
		interestPart := balance * monthlyRate
		principalPart := monthlyPayment - interestPart
		balance -= principalPart
		// Clamp tiny floating-point drift on the last payment
		if balance < 0 {
			balance = 0
		}
		schedule = append(schedule, AmortizationRow{
			Month:     m,
			Payment:   monthlyPayment,
			Principal: principalPart,
			Interest:  interestPart,
			Balance:   balance,
		})
	}

	totalPayment := monthlyPayment * float64(months)
	return LoanResult{
		Principal:      principal,
		AnnualRate:     annualRate,
		Years:          years,
		MonthlyPayment: monthlyPayment,
		TotalPayment:   totalPayment,
		TotalInterest:  totalPayment - principal,
		Schedule:       schedule,
	}
}
