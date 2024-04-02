package handler

import (
	"log"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/spf13/cobra"
)

// AdjustCmd represents the adjust command
var AdjustCmd = &cobra.Command{
	Use:   "adjust",
	Short: "Adjust the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		oldCategory, _ := cmd.Flags().GetString("oldcategory")
		newCategory, _ := cmd.Flags().GetString("newcategory")
		amount, _ := cmd.Flags().GetString("amount")
		err := budget_db.UpdateBudget(oldCategory, newCategory, amount, "")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	AdjustCmd.Flags().StringP("oldcategory", "o", "", "Write the old category name to adjust")
	AdjustCmd.Flags().StringP("newcategory", "n", "", "Write the new category name to allocate")
	AdjustCmd.Flags().StringP("amount", "a", "", "Write the new amount of the category to adjust")
}
