package total_amount_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	usecases_total_amount "github.com/ibilalkayy/flow/usecases/app/total_amount"
	"github.com/spf13/cobra"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		category, _ := cmd.Flags().GetString("category")
		label, _ := cmd.Flags().GetString("label")
		var c conversion.MyConversion
		totalAmount := c.StringToInt(amount)

		var set usecases_total_amount.MyTotalAmount
		err := set.SetTotalAmount(totalAmount, category, label)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	AddCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to add")
	AddCmd.Flags().StringP("category", "c", "", "Specify a category to include in the total amount")
	AddCmd.Flags().StringP("label", "l", "", "Provide a label for setting up your total amount. Write label b/w commas")
}
