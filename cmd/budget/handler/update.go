package budget_handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		oldCategory, _ := cmd.Flags().GetString("old-category")
		newCategory, _ := cmd.Flags().GetString("new-category")
		amount, _ := cmd.Flags().GetString("amount")

		h := TakeHandler()
		newAmount := h.Deps.Common.StringToInt(amount)

		bv := entities.BudgetVariables{
			Category: oldCategory,
			Amount:   newAmount,
		}
		err := h.Deps.ManageBudget.UpdateBudget(&bv, newCategory)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Your budget category is successfully updated!")
	},
}

func init() {
	UpdateCmd.Flags().StringP("old-category", "o", "", "Write the old category name to update")
	UpdateCmd.Flags().StringP("new-category", "n", "", "Write the new category name to allocate")
	UpdateCmd.Flags().StringP("amount", "a", "", "Write the new amount of the category to update")
}
