package budget_handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/spf13/cobra"
)

// ViewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		var m budget_db.MyBudgetDatabase
		details, err := m.ViewBudget(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(details[0])
	},
}

func init() {
	ViewCmd.Flags().StringP("category", "c", "", "Write the category name to show the specific details")
}
