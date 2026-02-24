package calculator

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RenderLoan(result LoanResult, showSchedule bool) string {
	p := message.NewPrinter(language.English)

	labelStyle := lipgloss.NewStyle().Bold(true).Align(lipgloss.Right).Width(18).Render
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render
	warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Render

	summary := fmt.Sprintf(
		"%s %s\n%s %s\n%s %s\n%s %s\n%s %s\n%s %s",
		labelStyle("Principal:"), valueStyle(p.Sprintf("$%.2f", result.Principal)),
		labelStyle("Annual Rate:"), valueStyle(p.Sprintf("%.2f%%", result.AnnualRate)),
		labelStyle("Term:"), valueStyle(p.Sprintf("%d years (%d months)", result.Years, result.Years*12)),
		labelStyle("Monthly Payment:"), valueStyle(p.Sprintf("$%.2f", result.MonthlyPayment)),
		labelStyle("Total Payment:"), valueStyle(p.Sprintf("$%.2f", result.TotalPayment)),
		labelStyle("Total Interest:"), warnStyle(p.Sprintf("$%.2f", result.TotalInterest)),
	)

	if !showSchedule {
		return summary
	}

	// Build amortization table (yearly summary rows)
	headers := []string{"MONTH", "PAYMENT", "PRINCIPAL", "INTEREST", "BALANCE"}
	var rows [][]string
	for _, row := range result.Schedule {
		rows = append(rows, []string{
			p.Sprintf("%d", row.Month),
			p.Sprintf("$%.2f", row.Payment),
			p.Sprintf("$%.2f", row.Principal),
			p.Sprintf("$%.2f", row.Interest),
			p.Sprintf("$%.2f", row.Balance),
		})
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
				return lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Padding(0, 1).Align(lipgloss.Right)
			case 3:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Padding(0, 1).Align(lipgloss.Right)
			case 4:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Padding(0, 1).Align(lipgloss.Right)
			default:
				return lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Right)
			}
		})

	return summary + "\n\n" + t.String()
}
