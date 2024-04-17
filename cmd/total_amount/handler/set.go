package total_amount_handler

import (
	"log"

	internal_total_amount "github.com/ibilalkayy/flow/internal/app/total_amount"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/spf13/cobra"
)

// SetCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		include_category, _ := cmd.Flags().GetString("include")
		label, _ := cmd.Flags().GetString("label")
		totalAmount := functions.StringToInt(amount)

		err := internal_total_amount.SetTotalAmount(totalAmount, include_category, label)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	SetCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to set")
	SetCmd.Flags().StringP("include", "i", "", "Specify a category to include in the total amount")
	SetCmd.Flags().StringP("label", "l", "", "Provide a label for setting up your total amount. Write label b/w commas")
}
