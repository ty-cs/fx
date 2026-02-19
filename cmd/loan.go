package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ty-cs/fx/pkg/calculator"
	"github.com/ty-cs/fx/util"
)

var (
	loanPrincipal util.MoneyValue
	loanRate      float64
	loanYears     int
	loanSchedule  bool
)

var loanCmd = &cobra.Command{
	Use:   "loan",
	Short: "Calculates loan or mortgage payments and amortization.",
	Long: `Calculates monthly payments, total interest, and total cost
for a loan or mortgage given a principal, annual interest rate, and term.

Use --schedule to print the full month-by-month amortization table.

Example:
  fx loan --principal 500k --rate 6.5 --years 30
  fx loan --principal 25000 --rate 4.9 --years 5 --schedule`,
	Run: func(cmd *cobra.Command, args []string) {
		result := calculator.CalculateLoan(float64(loanPrincipal), loanRate, loanYears)
		fmt.Println(calculator.RenderLoan(result, loanSchedule))
	},
}

func init() {
	rootCmd.AddCommand(loanCmd)

	loanCmd.Flags().VarP(&loanPrincipal, "principal", "p", "Loan amount (required, supports shorthand: 500k, 1m)")
	loanCmd.Flags().Float64VarP(&loanRate, "rate", "r", 0, "Annual interest rate in percent (e.g., 6.5 for 6.5%)")
	loanCmd.Flags().IntVarP(&loanYears, "years", "y", 0, "Loan term in years (required)")
	loanCmd.Flags().BoolVarP(&loanSchedule, "schedule", "s", false, "Print the full month-by-month amortization schedule")

	if err := loanCmd.MarkFlagRequired("principal"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := loanCmd.MarkFlagRequired("years"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := loanCmd.MarkFlagRequired("rate"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
