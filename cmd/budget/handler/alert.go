package handler

import (
	"log"

	app "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/structs"
	"github.com/spf13/cobra"
)

// AlertCmd represents the alert command
var AlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Get notification once you pass the budget",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		frequency, _ := cmd.Flags().GetString("frequency")
		method, _ := cmd.Flags().GetString("method")

		av := structs.AlertVariables{
			Category:  category,
			Frequency: frequency,
			Method:    method,
		}

		err := app.AlertSetup(av)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	AlertCmd.AddCommand(MsgCmd)
	AlertCmd.Flags().StringP("category", "c", "", "Write the category name to take its budget amount")
	AlertCmd.Flags().StringP("frequency", "f", "", "Write the frequency of notifications (e.g., hourly, daily, weekly, monthly)")
	AlertCmd.Flags().StringP("method", "m", "", "Write the preferred method of notification [email or CLI] message")
}
