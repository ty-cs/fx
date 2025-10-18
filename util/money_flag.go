package util

import (
	"fmt"
	"strconv"
	"strings"
)

// MoneyValue is a custom flag type for handling money shorthands (e.g., 10k, 1M).
type MoneyValue float64

// String returns the string representation of the MoneyValue.
func (m *MoneyValue) String() string {
	return fmt.Sprintf("%f", *m)
}

// Set parses the string value and sets the MoneyValue.
func (m *MoneyValue) Set(s string) error {
	s = strings.ToLower(strings.TrimSpace(s))
	multiplier := 1.0
	if strings.HasSuffix(s, "k") {
		multiplier = 1e3
		s = strings.TrimSuffix(s, "k")
	} else if strings.HasSuffix(s, "m") {
		multiplier = 1e6
		s = strings.TrimSuffix(s, "m")
	} else if strings.HasSuffix(s, "b") {
		multiplier = 1e9
		s = strings.TrimSuffix(s, "b")
	}

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid money value: %s", s)
	}

	*m = MoneyValue(val * multiplier)
	return nil
}

// Type returns the type of the flag.
func (m *MoneyValue) Type() string {
	return "moneyValue"
}
