package spend

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/cmd"
	spend_handler "github.com/ibilalkayy/flow/cmd/spend/handler"
	conversion "github.com/ibilalkayy/flow/common"
	usecases_spending "github.com/ibilalkayy/flow/usecases/app/spend"
	"github.com/spf13/cobra"
)

// spendCmd represents the spend command
var SpendCmd = &cobra.Command{
	Use:   "spend",
	Short: "Spending money on various categories",
	Run: func(cmd *cobra.Command, args []string) {
		categoryName, _ := cmd.Flags().GetString("category")
		spendingAmount, _ := cmd.Flags().GetString("amount")
		var c conversion.MyConversion
		spendingAmountInt := c.StringToInt(spendingAmount)

		if len(categoryName) != 0 && spendingAmountInt != 0 {
			var m usecases_spending.MySpending
			err := m.SpendMoney(categoryName, spendingAmountInt)
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
	// Subcommand
	SpendCmd.AddCommand(spend_handler.HistoryCmd)

	// Flags
	SpendCmd.Flags().StringP("category", "c", "", "Write the category name to spend the money on")
	SpendCmd.Flags().StringP("amount", "a", "", "Write the spending amount for a category")
}
