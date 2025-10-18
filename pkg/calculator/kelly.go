package calculator

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type KellyResult struct {
	WinProbability float64
	NetOdds        float64
	KellyFraction  float64
}

func CalculateKelly(winProbability, netOdds float64) KellyResult {
	lossProbability := 1 - winProbability
	kellyFraction := (netOdds*winProbability - lossProbability) / netOdds

	return KellyResult{
		WinProbability: winProbability,
		NetOdds:        netOdds,
		KellyFraction:  kellyFraction,
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

	summary := RenderKellySummary(result.KellyFraction)

	return fmt.Sprintf("%s\n%s\n%s", winProb, netOdds, kellyFraction) + "\n\n" + summary
}

func RenderKellySummary(kellyFraction float64) string {
	if kellyFraction > 0 {
		style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))
		return style.Render(fmt.Sprintf("Recommended bet size is %.2f%% of your bankroll.", kellyFraction*100))
	}
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("3"))
	return style.Render("The Kelly Criterion suggests not to bet.")
}
