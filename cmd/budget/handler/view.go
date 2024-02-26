package handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/internal/app"
	"github.com/spf13/cobra"
)

// ViewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		details, err := app.ViewBudget(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(details)
	},
}

func init() {
	ViewCmd.Flags().StringP("category", "c", "", "Write the category name to show the specific details")
}
