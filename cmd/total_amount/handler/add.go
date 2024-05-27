package total_amount_handler

import (
	"log"

	"github.com/spf13/cobra"
)

// AddCmd represents the set command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		category, _ := cmd.Flags().GetString("category")
		label, _ := cmd.Flags().GetString("label")

		h := TakeHandler()
		totalAmount := h.Deps.Common.StringToInt(amount)
		err := h.Deps.Total.SetTotalAmount(totalAmount, category, label)
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
