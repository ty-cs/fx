package calculator

import (
	"fmt"
	"math"
	"testing"
)

// tolerance for floating-point comparisons
const epsilon = 0.01

func TestCalculateLoan_MonthlyPayment(t *testing.T) {
	tests := []struct {
		name              string
		principal         float64
		annualRate        float64
		years             int
		wantMonthly       float64
		wantTotalInterest float64
	}{
		{
			name:              "30yr_mortgage_6.5pct",
			principal:         500_000,
			annualRate:        6.5,
			years:             30,
			wantMonthly:       3160.34,
			wantTotalInterest: 637_722.44,
		},
		{
			name:              "5yr_car_loan_4.9pct",
			principal:         25_000,
			annualRate:        4.9,
			years:             5,
			wantMonthly:       470.64,
			wantTotalInterest: 3238.18,
		},
		{
			name:              "10yr_personal_loan_8pct",
			principal:         50_000,
			annualRate:        8.0,
			years:             10,
			wantMonthly:       606.64,
			wantTotalInterest: 22_796.80,
		},
		{
			name:              "zero_interest_loan",
			principal:         12_000,
			annualRate:        0,
			years:             1,
			wantMonthly:       1_000.00,
			wantTotalInterest: 0.00,
		},
		{
			name:              "1yr_high_rate_24pct",
			principal:         5_000,
			annualRate:        24.0,
			years:             1,
			wantMonthly:       472.80,
			wantTotalInterest: 673.58,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateLoan(tc.principal, tc.annualRate, tc.years)

			if math.Abs(result.MonthlyPayment-tc.wantMonthly) > epsilon {
				t.Errorf("MonthlyPayment: got %.2f, want %.2f", result.MonthlyPayment, tc.wantMonthly)
			}
			if math.Abs(result.TotalInterest-tc.wantTotalInterest) > 1.0 { // larger tolerance for cumulative
				t.Errorf("TotalInterest: got %.2f, want %.2f", result.TotalInterest, tc.wantTotalInterest)
			}
		})
	}
}

func TestCalculateLoan_TotalPaymentIdentity(t *testing.T) {
	// TotalPayment must always equal Principal + TotalInterest.
	cases := []struct {
		principal, rate float64
		years           int
	}{
		{100_000, 5.0, 20},
		{250_000, 3.75, 15},
		{10_000, 0, 3},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%.0f_%.2f_%d", tc.principal, tc.rate, tc.years), func(t *testing.T) {
			r := CalculateLoan(tc.principal, tc.rate, tc.years)
			want := r.Principal + r.TotalInterest
			if math.Abs(r.TotalPayment-want) > epsilon {
				t.Errorf("TotalPayment %.2f != Principal+TotalInterest %.2f", r.TotalPayment, want)
			}
		})
	}
}

func TestCalculateLoan_AmortizationScheduleLength(t *testing.T) {
	tests := []struct{ years, wantMonths int }{
		{1, 12},
		{5, 60},
		{30, 360},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d_years", tc.years), func(t *testing.T) {
			r := CalculateLoan(100_000, 5.0, tc.years)
			if len(r.Schedule) != tc.wantMonths {
				t.Errorf("Schedule length: got %d, want %d", len(r.Schedule), tc.wantMonths)
			}
		})
	}
}

func TestCalculateLoan_ScheduleFinalBalance(t *testing.T) {
	// The last row's Balance should be ≈ 0 (fully paid off).
	cases := []struct {
		principal, rate float64
		years           int
	}{
		{300_000, 7.0, 30},
		{15_000, 5.5, 3},
		{50_000, 0, 5},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%.0f_%.2f_%d", tc.principal, tc.rate, tc.years), func(t *testing.T) {
			r := CalculateLoan(tc.principal, tc.rate, tc.years)
			last := r.Schedule[len(r.Schedule)-1]
			if math.Abs(last.Balance) > 0.05 {
				t.Errorf("Final balance should be ~0, got %.4f", last.Balance)
			}
		})
	}
}

func TestCalculateLoan_ScheduleEachRowPrincipalPlusInterestEqualsPayment(t *testing.T) {
	// For every row: Payment ≈ PrincipalPart + InterestPart
	r := CalculateLoan(200_000, 6.0, 25)
	for _, row := range r.Schedule {
		diff := math.Abs(row.Payment - (row.Principal + row.Interest))
		if diff > epsilon {
			t.Errorf("Month %d: payment %.4f != principal %.4f + interest %.4f",
				row.Month, row.Payment, row.Principal, row.Interest)
		}
	}
}

func TestCalculateLoan_EarlyRowsHaveMoreInterest(t *testing.T) {
	// In a standard amortizing loan, interest portion should be higher at the
	// start than the end.
	r := CalculateLoan(400_000, 5.5, 30)
	first := r.Schedule[0]
	last := r.Schedule[len(r.Schedule)-1]
	if first.Interest <= last.Interest {
		t.Errorf("Interest should decrease over time: month 1=%.2f, last month=%.2f",
			first.Interest, last.Interest)
	}
}
