package calculator

import "math"

// PV calculates the present value of a future sum of money.
func PV(futureValue, rate float64, years int) float64 {
	return futureValue / math.Pow(1+rate, float64(years))
}
