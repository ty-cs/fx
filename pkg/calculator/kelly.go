package calculator

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type KellyResult struct {
	WinProbability    float64
	NetOdds           float64
	KellyFraction     float64
	HalfKellyFraction float64
}

func CalculateKelly(winProbability, netOdds float64) KellyResult {
	lossProbability := 1 - winProbability
	kellyFraction := (netOdds*winProbability - lossProbability) / netOdds

	return KellyResult{
		WinProbability:    winProbability,
		NetOdds:           netOdds,
		KellyFraction:     kellyFraction,
		HalfKellyFraction: kellyFraction / 2,
	}
}

func RenderKelly(result KellyResult) string {
	var (
		keyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Bold(true).Width(16)
		valStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	)

	winProb := fmt.Sprintf("%s %s", keyStyle.Render("Win Probability:"), valStyle.Render(fmt.Sprintf("%.2f", result.WinProbability)))
	netOdds := fmt.Sprintf("%s %s", keyStyle.Render("Net Odds:"), valStyle.Render(fmt.Sprintf("%.2f", result.NetOdds)))
	kellyFraction := fmt.Sprintf("%s %s", keyStyle.Render("Kelly Fraction:"), valStyle.Render(fmt.Sprintf("%.2f", result.KellyFraction)))

	summary := RenderKellySummary(result)

	return fmt.Sprintf("%s\n%s\n%s", winProb, netOdds, kellyFraction) + "\n\n" + summary
}

func RenderKellySummary(result KellyResult) string {
	if result.KellyFraction > 0 {
		style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))
		summary := fmt.Sprintf("Recommended bet size is %.2f%% (Half-Kelly: %.2f%%) of your bankroll.", result.KellyFraction*100, result.HalfKellyFraction*100)
		return style.Render(summary)
	}
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("3"))
	return style.Render("The Kelly Criterion suggests not to bet.")
}
