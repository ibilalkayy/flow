package handler

import (
	"log"

	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/ibilalkayy/flow/internal/structs"
	"github.com/spf13/cobra"
)

// SetupCmd represents the setup command
var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup for alert notification",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		frequency, _ := cmd.Flags().GetString("frequency")
		method, _ := cmd.Flags().GetString("method")

		av := structs.AlertVariables{
			Category:  category,
			Frequency: frequency,
			Method:    method,
		}

		err := internal_alert.AlertSetup(&av)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	SetupCmd.Flags().StringP("category", "c", "", "Write the category name to take its budget amount")
	SetupCmd.Flags().StringP("frequency", "f", "", "Write the frequency of notifications (e.g., hourly, daily, weekly, monthly)")
	SetupCmd.Flags().StringP("method", "m", "", "Write the preferred method of notification [email or CLI] message")
}
