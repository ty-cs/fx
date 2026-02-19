package calculator

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func RenderRule72(result Rule72Result) string {
	labelStyle := lipgloss.NewStyle().Bold(true).Align(lipgloss.Right).Width(30).Render
	mainStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true).Render
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render

	var lines string
	if result.RateMode {
		lines = fmt.Sprintf(
			"%s %s\n%s %s\n%s %s",
			labelStyle("Annual Rate:"), dimStyle(fmt.Sprintf("%.2f%%", result.Rate)),
			labelStyle("Years to double (Rule 72):"), mainStyle(fmt.Sprintf("%.1f years", result.Years)),
			labelStyle("Years to double (exact):"), dimStyle(fmt.Sprintf("%.2f years", result.ExactYears)),
		)
	} else {
		lines = fmt.Sprintf(
			"%s %s\n%s %s\n%s %s",
			labelStyle("Years to double:"), dimStyle(fmt.Sprintf("%.1f years", result.Years)),
			labelStyle("Rate needed (Rule 72):"), mainStyle(fmt.Sprintf("%.2f%%", result.Rate)),
			labelStyle("Rate needed (exact):"), dimStyle(fmt.Sprintf("%.2f%%", result.ExactRate)),
		)
	}

	return lines
}
