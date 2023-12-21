/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/piotrostr/aoc23/four"
	"github.com/spf13/cobra"
)

// fourCmd represents the four command
var fourCmd = &cobra.Command{
	Use:   "four",
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

		res, err := four.Scratchcards(string(input))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(*res)
	},
}

func init() {
	rootCmd.AddCommand(fourCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fourCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fourCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
