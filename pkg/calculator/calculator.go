package calculator

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type YearResult struct {
	Year           int
	Contributions  float64
	InterestEarned float64
	EndBalance     float64
	RealValue      float64 // inflation-adjusted value; 0 when no inflation rate given
}

func Calculate(principal, contribution, growth float64, years int, rate, inflation float64) ([]YearResult, float64) {
	var results []YearResult
	balance := principal
	totalContributions := 0.0
	totalInterest := 0.0
	currentContribution := contribution

	// Real rate of return using Fisher equation: (1+nominal)/(1+inflation) - 1
	// This properly accounts for contributions made at different points in time.
	realRate := 0.0
	if inflation > 0 {
		realRate = (1+rate/100)/(1+inflation/100) - 1
	}
	realBalance := principal

	for i := 1; i <= years; i++ {
		totalContributions += currentContribution
		interest := balance * (rate / 100)
		totalInterest += interest
		balance += interest
		// Add the current year's contribution at the end of the year
		balance += currentContribution

		// Real (inflation-adjusted) value using real rate of return.
		// Contributions are kept in today's dollars; only the growth rate is adjusted.
		realValue := 0.0
		if inflation > 0 {
			realBalance = realBalance*(1+realRate) + currentContribution
			realValue = realBalance
		}

		results = append(results, YearResult{
			Year:           i,
			Contributions:  totalContributions,
			InterestEarned: totalInterest,
			EndBalance:     balance,
			RealValue:      realValue,
		})

		currentContribution *= 1 + (growth / 100)
	}

	return results, balance
}

func RenderTable(results []YearResult, inflation float64) string {
	p := message.NewPrinter(language.English)
	showReal := inflation > 0

	headers := []string{"YEAR", "END BALANCE", "RETURN", "CONTRIBUTIONS"}
	if showReal {
		headers = append(headers, "REAL VALUE")
	}

	var rows [][]string
	for _, r := range results {
		row := []string{
			p.Sprintf("%d", r.Year),
			p.Sprintf("$%.2f", r.EndBalance),
			p.Sprintf("$%.2f", r.InterestEarned),
			p.Sprintf("$%.2f", r.Contributions),
		}
		if showReal {
			row = append(row, p.Sprintf("$%.2f", r.RealValue))
		}
		rows = append(rows, row)
	}

	numCols := 4
	if showReal {
		numCols = 5
	}

	t := table.New().
		Headers(headers...).
		Rows(rows...).
		Border(lipgloss.ASCIIBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("7"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch col {
			case 1:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Padding(0, 1).Align(lipgloss.Right)
			case 2:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Padding(0, 1).Align(lipgloss.Right)
			case 3:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Padding(0, 1).Align(lipgloss.Right)
			case 4:
				if numCols == 5 {
					return lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Padding(0, 1).Align(lipgloss.Right)
				}
				return lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Right)
			default:
				return lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Right)
			}
		})

	return t.String()
}

func RenderSummary(finalBalance float64, years int) string {
	p := message.NewPrinter(language.English)
	cyan := lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Render
	return p.Sprintf("\nFinal Investment Value after %d years: %s\n", years, cyan(p.Sprintf("$%.2f", finalBalance)))
}
