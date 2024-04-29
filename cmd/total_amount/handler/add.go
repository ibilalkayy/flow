package total_amount_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common/utils"
	usecases_total_amount "github.com/ibilalkayy/flow/usecases/app/total_amount"
	"github.com/spf13/cobra"
)

// AddCmd represents the set command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		include_category, _ := cmd.Flags().GetString("include")
		label, _ := cmd.Flags().GetString("label")
		totalAmount := conversion.StringToInt(amount)

		err := usecases_total_amount.SetTotalAmount(totalAmount, include_category, label)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	AddCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to add")
	AddCmd.Flags().StringP("include", "i", "", "Specify a category to include in the total amount")
	AddCmd.Flags().StringP("label", "l", "", "Provide a label for setting up your total amount. Write label b/w commas")
}
