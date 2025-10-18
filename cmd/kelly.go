/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ty-cs/fx/pkg/calculator"
)

var (
	winProbability float64
	netOdds        float64
)

// kellyCmd represents the kelly command
var kellyCmd = &cobra.Command{
	Use:   "kelly",
	Short: "Calculates the optimal bet size using the Kelly Criterion",
	Long: `Calculates the optimal bet size based on the Kelly Criterion formula.

The Kelly Criterion is a formula used to determine the optimal theoretical
size for a bet. It is f* = (bp - q) / b where:
- f* is the fraction of the current bankroll to wager.
- b is the net odds received on the wager (e.g., 0.2 for 20% return).
- p is the probability of winning.
- q is the probability of losing (1 - p).

Example:
fx kelly --win-probability 0.6 --net-odds 0.2
`,
	Run: func(cmd *cobra.Command, args []string) {
		if winProbability <= 0 || winProbability >= 1 {
			fmt.Println("Error: Win probability must be between 0 and 1.")
			os.Exit(1)
		}
		if netOdds <= 0 {
			fmt.Println("Error: Net odds must be greater than 0.")
			os.Exit(1)
		}

		result := calculator.CalculateKelly(winProbability, netOdds)
		fmt.Println(calculator.RenderKelly(result))
	},
}

func init() {
	rootCmd.AddCommand(kellyCmd)

	kellyCmd.Flags().Float64VarP(&winProbability, "win-probability", "p", 0, "Win probability (e.g., 0.6 for 60%)")
	kellyCmd.Flags().Float64VarP(&netOdds, "net-odds", "o", 0, "Net odds (e.g., 0.2 for 20% return)")
	_ = kellyCmd.MarkFlagRequired("win-probability")
	_ = kellyCmd.MarkFlagRequired("net-odds")
}
