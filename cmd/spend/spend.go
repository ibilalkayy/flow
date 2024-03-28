package spend

import (
	"log"

	"github.com/ibilalkayy/flow/cmd"
	internal_spending "github.com/ibilalkayy/flow/internal/app/spend"
	"github.com/spf13/cobra"
)

// spendCmd represents the spend command
var SpendCmd = &cobra.Command{
	Use:   "spend",
	Short: "Spending money on various categories",
	Run: func(cmd *cobra.Command, args []string) {
		categoryName, _ := cmd.Flags().GetString("category")
		spendingAmount, _ := cmd.Flags().GetString("amount")
		err := internal_spending.SpendMoney(categoryName, spendingAmount)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(SpendCmd)
	SpendCmd.Flags().StringP("category", "c", "", "Write the category name to spend the money on")
	SpendCmd.Flags().StringP("amount", "a", "", "Write the spending amount for a category")
}
