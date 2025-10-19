package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/ty-cs/fx/pkg/calculator"
	"github.com/ty-cs/fx/util"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/spf13/cobra"
)

var (
	futureValue util.MoneyValue
)

// pvCmd represents the pv command
var pvCmd = &cobra.Command{
	Use:   "pv",
	Short: "Calculates the present value of a future sum of money.",
	Long:  `Calculates the present value of a future sum of money, given a specific discount rate.`,
	Run: func(cmd *cobra.Command, args []string) {
		presentValue := calculator.PV(float64(futureValue), rate/100, years)
		p := message.NewPrinter(language.English)

		// Styles
		labelStyle := lipgloss.NewStyle().Bold(true).Align(lipgloss.Right).Width(14).Render
		valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render

		// Output
		fmt.Println(lipgloss.NewStyle().Render(
			fmt.Sprintf("%s %s\n%s %s\n%s %s\n\n%s %s",
				labelStyle("Future Value:"), valueStyle(p.Sprintf("$%.2f", float64(futureValue))),
				labelStyle("Discount Rate:"), valueStyle(p.Sprintf("%.2f%%", rate)),
				labelStyle("Years:"), valueStyle(p.Sprintf("%d", years)),
				labelStyle("Present Value:"), valueStyle(p.Sprintf("$%.2f", presentValue)),
			),
		))
	},
}

func init() {
	rootCmd.AddCommand(pvCmd)

	pvCmd.Flags().VarP(&futureValue, "fv", "f", "The future value of the money (required)")
	pvCmd.Flags().IntVarP(&years, "years", "y", 0, "The total duration in years (required)")
	pvCmd.Flags().Float64VarP(&rate, "rate", "r", 0, "The discount rate (e.g., enter 8 for 8%)")

	if err := pvCmd.MarkFlagRequired("fv"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := pvCmd.MarkFlagRequired("years"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := pvCmd.MarkFlagRequired("rate"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
