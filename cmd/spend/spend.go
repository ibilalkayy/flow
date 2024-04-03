package spend

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/cmd"
	internal_spending "github.com/ibilalkayy/flow/internal/app/spend"
	"github.com/ibilalkayy/flow/internal/structs"
	"github.com/spf13/cobra"
)

// spendCmd represents the spend command
var SpendCmd = &cobra.Command{
	Use:   "spend",
	Short: "Spending money on various categories",
	Run: func(cmd *cobra.Command, args []string) {
		categoryName, _ := cmd.Flags().GetString("category")
		spendingAmount, _ := cmd.Flags().GetString("amount")
		spendingAmountInt := structs.StringToInt(spendingAmount)

		if len(categoryName) != 0 && spendingAmountInt != 0 {
			err := internal_spending.SpendMoney(categoryName, spendingAmountInt)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(errors.New("select the command and flags"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(SpendCmd)
	SpendCmd.Flags().StringP("category", "c", "", "Write the category name to spend the money on")
	SpendCmd.Flags().StringP("amount", "a", "", "Write the spending amount for a category")
}
