/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/piotrostr/aoc23/three"
	"github.com/spf13/cobra"
)

// threeCmd represents the three command
var threeCmd = &cobra.Command{
	Use:   "three",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, err := os.ReadFile("three/input.txt")
		if err != nil {
			log.Fatal(err)
		}

		res, err := three.GearRatios(string(input))
		if err != nil {
			log.Fatal(err)
		}
		println(*res)
	},
}

func init() {
	rootCmd.AddCommand(threeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// threeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// threeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
