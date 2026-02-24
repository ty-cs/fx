package calculator

import (
	"math"
	"testing"
)

func TestRule72FromRate_YearsApproximation(t *testing.T) {
	tests := []struct {
		name      string
		rate      float64
		wantYears float64 // Rule-of-72 approximation
		wantExact float64 // logarithmic exact value
	}{
		{"8pct_classic", 8.0, 9.0, 9.01},
		{"6pct", 6.0, 12.0, 11.90},
		{"10pct", 10.0, 7.2, 7.27},
		{"1pct_low_rate", 1.0, 72.0, 69.66},
		{"36pct_high_rate", 36.0, 2.0, 2.25},
		{"12pct", 12.0, 6.0, 6.12},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := Rule72FromRate(tc.rate)

			if r.Rate != tc.rate {
				t.Errorf("Rate: got %.2f, want %.2f", r.Rate, tc.rate)
			}
			if !r.RateMode {
				t.Error("RateMode should be true")
			}
			// Rule-of-72 approximation: within 0.1 years
			if math.Abs(r.Years-tc.wantYears) > 0.1 {
				t.Errorf("Years (Rule72): got %.2f, want %.2f", r.Years, tc.wantYears)
			}
			// Exact value: within 0.05 years
			if math.Abs(r.ExactYears-tc.wantExact) > 0.05 {
				t.Errorf("ExactYears: got %.2f, want %.2f", r.ExactYears, tc.wantExact)
			}
		})
	}
}

func TestRule72FromRate_Formula(t *testing.T) {
	// Verify: Years == 72 / Rate
	rates := []float64{5, 7, 9, 15, 24}
	for _, rate := range rates {
		r := Rule72FromRate(rate)
		want := 72.0 / rate
		if math.Abs(r.Years-want) > 1e-9 {
			t.Errorf("rate=%.0f: got %.6f, want %.6f", rate, r.Years, want)
		}
	}
}

func TestRule72FromYears_RateApproximation(t *testing.T) {
	tests := []struct {
		name      string
		years     float64
		wantRate  float64 // Rule-of-72 approximation
		wantExact float64 // exact rate
	}{
		{"10_years", 10.0, 7.20, 7.18},
		{"9_years", 9.0, 8.00, 8.01},
		{"6_years", 6.0, 12.00, 12.25},
		{"20_years", 20.0, 3.60, 3.53},
		{"5_years", 5.0, 14.40, 14.87},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := Rule72FromYears(tc.years)

			if r.Years != tc.years {
				t.Errorf("Years: got %.2f, want %.2f", r.Years, tc.years)
			}
			if r.RateMode {
				t.Error("RateMode should be false")
			}
			// Rule-of-72 approximation: within 0.1%
			if math.Abs(r.Rate-tc.wantRate) > 0.1 {
				t.Errorf("Rate (Rule72): got %.2f, want %.2f", r.Rate, tc.wantRate)
			}
			// Exact value: within 0.05%
			if math.Abs(r.ExactRate-tc.wantExact) > 0.05 {
				t.Errorf("ExactRate: got %.2f, want %.2f", r.ExactRate, tc.wantExact)
			}
		})
	}
}

func TestRule72FromYears_Formula(t *testing.T) {
	// Verify: Rate == 72 / Years
	years := []float64{8, 12, 15, 18, 24}
	for _, y := range years {
		r := Rule72FromYears(y)
		want := 72.0 / y
		if math.Abs(r.Rate-want) > 1e-9 {
			t.Errorf("years=%.0f: got %.6f, want %.6f", y, r.Rate, want)
		}
	}
}

func TestRule72_Inverse(t *testing.T) {
	// Rule72FromRate(r).Years and Rule72FromYears(y).Rate should be inverses.
	rates := []float64{4, 6, 8, 10, 12}
	for _, rate := range rates {
		from_rate := Rule72FromRate(rate)
		from_years := Rule72FromYears(from_rate.Years)
		if math.Abs(from_years.Rate-rate) > 1e-9 {
			t.Errorf("rate=%.0f: inverse mismatch, got %.6f", rate, from_years.Rate)
		}
	}
}

func TestRule72_ExactDoublingVerification(t *testing.T) {
	// Verify exact years: (1 + rate/100)^exactYears == 2.0 (within epsilon)
	rates := []float64{5, 7, 10, 15}
	for _, rate := range rates {
		r := Rule72FromRate(rate)
		got := math.Pow(1+rate/100, r.ExactYears)
		if math.Abs(got-2.0) > 1e-6 {
			t.Errorf("rate=%.0f: (1+r)^exactYears=%.8f, want 2.0", rate, got)
		}
	}
}
