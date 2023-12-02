/*
Copyright © 2023 piotrostr
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc23",
	Short: "aoc23 is a CLI for Advent of Code 2023",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("info", "t", false, "for higher log verbosity")
}
