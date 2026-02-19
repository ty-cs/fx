package calculator

import (
	"math"
	"testing"
)

func TestCalculate_BasicGrowth(t *testing.T) {
	tests := []struct {
		name         string
		principal    float64
		contribution float64
		growth       float64
		years        int
		rate         float64
		inflation    float64
		wantBalance  float64
	}{
		{
			name:         "10k_8pct_10yr_no_contributions",
			principal:    10_000,
			contribution: 0,
			growth:       0,
			years:        10,
			rate:         8.0,
			inflation:    0,
			wantBalance:  21_589.25,
		},
		{
			name:         "1k_10pct_5yr_no_contributions",
			principal:    1_000,
			contribution: 0,
			growth:       0,
			years:        5,
			rate:         10.0,
			inflation:    0,
			wantBalance:  1_610.51,
		},
		{
			name:         "zero_rate_no_growth",
			principal:    5_000,
			contribution: 0,
			growth:       0,
			years:        10,
			rate:         0.0,
			inflation:    0,
			wantBalance:  5_000.00,
		},
		{
			name:         "with_contribution_no_growth",
			principal:    10_000,
			contribution: 1_200,
			growth:       0,
			years:        5,
			rate:         7.0,
			inflation:    0,
			// Year 1: (10000 * 1.07) + 1200 = 11900
			// Year 2: (11900 * 1.07) + 1200 = 13933
			// Year 3: (13933 * 1.07) + 1200 = 16108.31
			// Year 4: (16108.31 * 1.07) + 1200 = 18435.89
			// Year 5: (18435.89 * 1.07) + 1200 = 20926.40
			wantBalance: 20_926.40,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			results, finalBalance := Calculate(tc.principal, tc.contribution, tc.growth, tc.years, tc.rate, tc.inflation)

			if len(results) != tc.years {
				t.Errorf("Expected %d results, got %d", tc.years, len(results))
			}
			if math.Abs(finalBalance-tc.wantBalance) > 0.05 {
				t.Errorf("FinalBalance: got %.2f, want %.2f", finalBalance, tc.wantBalance)
			}
		})
	}
}

func TestCalculate_YearNumbers(t *testing.T) {
	results, _ := Calculate(10_000, 0, 0, 5, 8, 0)
	for i, r := range results {
		if r.Year != i+1 {
			t.Errorf("Result[%d].Year: got %d, want %d", i, r.Year, i+1)
		}
	}
}

func TestCalculate_BalanceAlwaysIncreases_PositiveRate(t *testing.T) {
	results, _ := Calculate(10_000, 0, 0, 10, 5, 0)
	for i := 1; i < len(results); i++ {
		if results[i].EndBalance <= results[i-1].EndBalance {
			t.Errorf("Balance should increase year-over-year: year %d (%.2f) <= year %d (%.2f)",
				results[i].Year, results[i].EndBalance, results[i-1].Year, results[i-1].EndBalance)
		}
	}
}

func TestCalculate_ContributionGrowth(t *testing.T) {
	// With 10% contribution growth, each year's contribution should be 10% higher.
	// We can verify indirectly: total contributions grows faster than linearly.
	results, _ := Calculate(0, 1_000, 10, 3, 0, 0)
	// Year 1: contribution = 1000, total = 1000
	// Year 2: contribution = 1100, total = 2100
	// Year 3: contribution = 1210, total = 3310
	wantContributions := []float64{1_000, 2_100, 3_310}
	for i, want := range wantContributions {
		if math.Abs(results[i].Contributions-want) > 0.01 {
			t.Errorf("Year %d Contributions: got %.2f, want %.2f", i+1, results[i].Contributions, want)
		}
	}
}

func TestCalculate_InterestEarned_Cumulative(t *testing.T) {
	// With zero contributions, InterestEarned should be monotonically increasing.
	results, _ := Calculate(10_000, 0, 0, 5, 8, 0)
	for i := 1; i < len(results); i++ {
		if results[i].InterestEarned <= results[i-1].InterestEarned {
			t.Errorf("InterestEarned should be cumulative and growing: year %d=%.2f, year %d=%.2f",
				results[i].Year, results[i].InterestEarned, results[i-1].Year, results[i-1].InterestEarned)
		}
	}
}

func TestCalculate_FinalBalanceMatchesLastResult(t *testing.T) {
	results, finalBalance := Calculate(50_000, 500, 5, 15, 7, 0)
	last := results[len(results)-1]
	if math.Abs(last.EndBalance-finalBalance) > 1e-9 {
		t.Errorf("FinalBalance (%.2f) != last result EndBalance (%.2f)", finalBalance, last.EndBalance)
	}
}

func TestCalculate_Inflation_NoInflation(t *testing.T) {
	// When inflation=0, RealValue should be 0 for all rows.
	results, _ := Calculate(10_000, 0, 0, 5, 8, 0)
	for _, r := range results {
		if r.RealValue != 0 {
			t.Errorf("Year %d: expected RealValue=0 when inflation=0, got %.2f", r.Year, r.RealValue)
		}
	}
}

func TestCalculate_Inflation_RealValueLessThanNominal(t *testing.T) {
	// With positive inflation, real value should always be less than nominal balance.
	results, _ := Calculate(10_000, 0, 0, 10, 8, 3)
	for _, r := range results {
		if r.RealValue >= r.EndBalance {
			t.Errorf("Year %d: RealValue (%.2f) should be < EndBalance (%.2f)", r.Year, r.RealValue, r.EndBalance)
		}
	}
}

func TestCalculate_Inflation_RealValueFormula(t *testing.T) {
	// RealValue = EndBalance / (1+inflation/100)^year
	results, _ := Calculate(10_000, 0, 0, 5, 8, 3)
	for _, r := range results {
		want := r.EndBalance / math.Pow(1.03, float64(r.Year))
		if math.Abs(r.RealValue-want) > 1e-6 {
			t.Errorf("Year %d: RealValue=%.6f, want %.6f", r.Year, r.RealValue, want)
		}
	}
}

func TestCalculate_Inflation_HighInflationLargelyErodesGains(t *testing.T) {
	// At 50% inflation, real value degrades dramatically year over year.
	results, _ := Calculate(10_000, 0, 0, 5, 8, 50)
	// Real value must strictly decrease even if nominal grows at 8%.
	for i := 1; i < len(results); i++ {
		if results[i].RealValue >= results[i-1].RealValue {
			t.Errorf("Year %d: RealValue (%.2f) should be less than year %d (%.2f) under 50pct inflation",
				results[i].Year, results[i].RealValue, results[i-1].Year, results[i-1].RealValue)
		}
	}
}

func TestCalculate_Inflation_WhenRateEqualsInflation_RealValueStable(t *testing.T) {
	// With zero contributions and rate == inflation, real value stays ≈ principal.
	results, _ := Calculate(10_000, 0, 0, 10, 5, 5)
	for _, r := range results {
		if math.Abs(r.RealValue-10_000) > 0.01 {
			t.Errorf("Year %d: expected RealValue≈10000 when rate==inflation, got %.2f", r.Year, r.RealValue)
		}
	}
}
