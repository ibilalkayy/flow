package budget_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the budget of different categories",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetString("amount")
		var c conversion.MyConversion
		amountInt := c.StringToInt(amount)

		bv := entities.BudgetVariables{Category: category, Amount: amountInt}
		var m budget_db.MyBudgetDatabase
		err := m.CreateBudget(&bv)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	CreateCmd.Flags().StringP("category", "c", "", "Write the category like groceries, utilities, etc")
	CreateCmd.Flags().StringP("amount", "a", "", "Write the total amount for that category")
}
