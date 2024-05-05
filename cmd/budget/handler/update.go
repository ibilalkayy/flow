package budget_handler

import (
	"fmt"
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		var c conversion.MyConversion
		var m budget_db.MyBudgetDatabase

		oldCategory, _ := cmd.Flags().GetString("oldcategory")
		newCategory, _ := cmd.Flags().GetString("newcategory")
		amount, _ := cmd.Flags().GetString("amount")
		newAmount := c.StringToInt(amount)

		err := m.UpdateBudget(oldCategory, newCategory, newAmount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Your budget category is successfully updated!")
	},
}

func init() {
	UpdateCmd.Flags().StringP("oldcategory", "o", "", "Write the old category name to update")
	UpdateCmd.Flags().StringP("newcategory", "n", "", "Write the new category name to allocate")
	UpdateCmd.Flags().StringP("amount", "a", "", "Write the new amount of the category to update")
}
