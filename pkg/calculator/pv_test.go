package calculator

import (
	"fmt"
	"math"
	"testing"
)

func TestPV(t *testing.T) {
	testCases := []struct {
		futureValue float64
		rate        float64
		years       int
		expected    float64
	}{
		{1000000, 0.07, 10, 508349.29},
		{25000, 0.05, 5, 19588.15},
		{0, 0.1, 20, 0},
		{1000, 0, 10, 1000},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("fv_%.2f_rate_%.2f_years_%d", tc.futureValue, tc.rate, tc.years), func(t *testing.T) {
			actual := PV(tc.futureValue, tc.rate, tc.years)
			if math.Abs(actual-tc.expected) > 0.01 {
				t.Errorf("Expected present value to be %.2f, but got %.2f", tc.expected, actual)
			}
		})
	}
}
