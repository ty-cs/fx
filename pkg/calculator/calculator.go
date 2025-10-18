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
}

func Calculate(principal, contribution, growth float64, years int, rate float64) ([]YearResult, float64) {
	var results []YearResult
	balance := principal
	totalContributions := 0.0
	totalInterest := 0.0
	currentContribution := contribution

	for i := 1; i <= years; i++ {
		totalContributions += currentContribution
		interest := balance * (rate / 100)
		totalInterest += interest
		balance += interest
		// Add the current year's contribution at the end of the year
		balance += currentContribution

		results = append(results, YearResult{
			Year:           i,
			Contributions:  totalContributions,
			InterestEarned: totalInterest,
			EndBalance:     balance,
		})

		currentContribution *= 1 + (growth / 100)
	}

	return results, balance
}

func RenderTable(results []YearResult) string {
	p := message.NewPrinter(language.English)

	headers := []string{"YEAR", "END BALANCE", "RETURN", "CONTRIBUTIONS"}
	var rows [][]string
	for _, r := range results {
		rows = append(rows, []string{
			p.Sprintf("%d", r.Year),
			p.Sprintf("$%.2f", r.EndBalance),
			p.Sprintf("$%.2f", r.InterestEarned),
			p.Sprintf("$%.2f", r.Contributions),
		})
	}

	t := table.New().
		Headers(headers...).
		Rows(rows...).
		Border(lipgloss.ASCIIBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("7"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case col == 1:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Padding(0, 1).Align(lipgloss.Right)
			case col == 2:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Padding(0, 1).Align(lipgloss.Right)
			case col == 3:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Padding(0, 1).Align(lipgloss.Right)
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
