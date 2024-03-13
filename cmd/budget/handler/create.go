package handler

import (
	"log"

	app "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the budget of different categories",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetString("amount")
		bv := &app.BudgetVariables{Category: category, Amount: amount}
		err := app.CreateBudget(bv, "db/budget_db/migrations/")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	CreateCmd.Flags().StringP("category", "c", "", "Write the category like groceries, utilities")
	CreateCmd.Flags().StringP("amount", "a", "", "Write the total amount for that category")
}
