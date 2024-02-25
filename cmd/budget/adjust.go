package budget

import (
	"log"

	"github.com/ibilalkayy/flow/internal/app"
	"github.com/spf13/cobra"
)

// adjustCmd represents the adjust command
var adjustCmd = &cobra.Command{
	Use:   "adjust",
	Short: "Adjust the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		oldCategory, _ := cmd.Flags().GetString("oldcategory")
		newCategory, _ := cmd.Flags().GetString("newcategory")
		amount, _ := cmd.Flags().GetString("amount")
		err := app.UpdateBudget(oldCategory, newCategory, amount)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	adjustCmd.Flags().StringP("oldcategory", "o", "", "Write the old category name to adjust")
	adjustCmd.Flags().StringP("newcategory", "n", "", "Write the new category name to allocate")
	adjustCmd.Flags().StringP("amount", "a", "", "Write the new amount of the category to adjust")
}
