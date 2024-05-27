package total_amount_handler

import (
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		old_category, _ := cmd.Flags().GetString("old-category")
		new_category, _ := cmd.Flags().GetString("new-category")
		amount, _ := cmd.Flags().GetString("amount")
		label, _ := cmd.Flags().GetString("label")

		h := TakeHandler()
		totalAmount := h.Deps.Common.StringToInt(amount)
		tv := entities.TotalAmountVariables{
			Included:    old_category,
			NewCategory: new_category,
			TotalAmount: totalAmount,
			Label:       label,
		}

		err := h.Deps.TotalAmount.UpdateTotalAmount(&tv)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UpdateCmd.Flags().StringP("old-category", "o", "", "Write the old category that you want to update")
	UpdateCmd.Flags().StringP("new-category", "n", "", "Write the new category to update with")
	UpdateCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to update")
	UpdateCmd.Flags().StringP("label", "l", "", "Write the label that you want to update")
}
