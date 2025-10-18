package cmd

import (
	"fmt"
	"os"

	"github.com/ty-cs/fx/pkg/calculator"
	"github.com/ty-cs/fx/util"

	"github.com/spf13/cobra"
)

var (
	principal    util.MoneyValue
	contribution util.MoneyValue
	growth       float64
	years        int
	rate         float64
)

// compoundCmd represents the compound command
var compoundCmd = &cobra.Command{
	Use:   "compound",
	Short: "Calculates compound growth for an investment.",
	Long: `Calculates the future value of an investment with compound growth
and displays a year-by-year breakdown of the growth in a formatted table.`,
	Run: func(cmd *cobra.Command, args []string) {
		results, finalBalance := calculator.Calculate(float64(principal), float64(contribution), growth, years, rate)
		fmt.Println(calculator.RenderTable(results))
		fmt.Println(calculator.RenderSummary(finalBalance, years))
	},
}

func init() {
	rootCmd.AddCommand(compoundCmd)

	compoundCmd.Flags().VarP(&principal, "principal", "p", "The initial investment principal (required)")
	compoundCmd.Flags().VarP(&contribution, "contribution", "c", "The annual contribution")
	compoundCmd.Flags().Float64VarP(&growth, "growth", "g", 0, "The annual percentage increase in the contribution (e.g., enter 3 for 3%)")
	compoundCmd.Flags().IntVarP(&years, "years", "y", 0, "The total investment duration in years (required)")
	compoundCmd.Flags().Float64VarP(&rate, "rate", "r", 0, "The expected annual rate of return (e.g., enter 8 for 8%)")

	if err := compoundCmd.MarkFlagRequired("principal"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := compoundCmd.MarkFlagRequired("years"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := compoundCmd.MarkFlagRequired("rate"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
