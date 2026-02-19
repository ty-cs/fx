package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ty-cs/fx/pkg/calculator"
)

var (
	r72Rate  float64
	r72Years float64
)

var rule72Cmd = &cobra.Command{
	Use:   "rule72",
	Short: "Estimates years to double (or rate needed) using the Rule of 72.",
	Long: `Estimates how long it takes an investment to double using the Rule of 72.

Provide either --rate to find the years to double, or --years to find the
required annual rate. Exactly one of these flags must be provided.

The Rule of 72: divide 72 by the annual rate to get approximate doubling time.
The exact value (using logarithms) is also shown for comparison.

Examples:
  fx rule72 --rate 8        # How many years to double at 8% annual return?
  fx rule72 --years 10      # What rate doubles money in 10 years?`,
	Run: func(cmd *cobra.Command, args []string) {
		rateSet := cmd.Flags().Changed("rate")
		yearsSet := cmd.Flags().Changed("years")

		if rateSet == yearsSet {
			fmt.Println("Error: provide exactly one of --rate or --years.")
			os.Exit(1)
		}

		var result calculator.Rule72Result
		if rateSet {
			if r72Rate <= 0 {
				fmt.Println("Error: --rate must be greater than 0.")
				os.Exit(1)
			}
			result = calculator.Rule72FromRate(r72Rate)
		} else {
			if r72Years <= 0 {
				fmt.Println("Error: --years must be greater than 0.")
				os.Exit(1)
			}
			result = calculator.Rule72FromYears(r72Years)
		}

		fmt.Println(calculator.RenderRule72(result))
	},
}

func init() {
	rootCmd.AddCommand(rule72Cmd)

	rule72Cmd.Flags().Float64VarP(&r72Rate, "rate", "r", 0, "Annual rate of return in percent (e.g., 8 for 8%)")
	rule72Cmd.Flags().Float64VarP(&r72Years, "years", "y", 0, "Desired years to double (e.g., 10)")
}
