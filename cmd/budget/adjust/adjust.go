package adjust

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adjustCmd represents the adjust command
var AdjustCmd = &cobra.Command{
	Use:   "adjust",
	Short: "Adjust the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		oldCategory, _ := cmd.Flags().GetString("oldcategory")
		newCategory, _ := cmd.Flags().GetString("newcategory")
		amount, _ := cmd.Flags().GetString("amount")
		fmt.Println(oldCategory)
		fmt.Println(newCategory)
		fmt.Println(amount)
	},
}

func init() {
	AdjustCmd.Flags().StringP("oldcategory", "o", "", "Write the old category name to adjust")
	AdjustCmd.Flags().StringP("newcategory", "n", "", "Write the new category name to allocate")
	AdjustCmd.Flags().StringP("amount", "a", "", "Write the new amount of the category to adjust")
}
