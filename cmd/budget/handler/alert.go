package handler

import (
	"fmt"
	"log"

	app "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/spf13/cobra"
)

// alertCmd represents the alert command
var AlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Get notification once you pass the budget",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		frequency, _ := cmd.Flags().GetString("frequency")
		method, _ := cmd.Flags().GetString("method")
		amount, err := app.BudgetAmount(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(frequency)
		fmt.Println(method)
		fmt.Println(amount)
	},
}

func init() {
	AlertCmd.Flags().StringP("category", "c", "", "Write the category name to take it's budget amount")
	AlertCmd.Flags().StringP("frequency", "f", "", "Write the daily, weekly, or monthly frequency")
	AlertCmd.Flags().StringP("method", "m", "", "Write the preferred method of notification either email or CLI messsage")
}
