/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.1.46"

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "flow",
	Version: version,
	Short:   "Manage budgets with Formance API",
	Long: `Flow is a budget planning application designed to empower users with the ability to track, 
analyze, and optimize their spending habits and financial goals. With a user-friendly CLI. 

It manages the finances and achieve greater financial stability  by leveraging
the Formance API for a comprehensive financial management solution.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
