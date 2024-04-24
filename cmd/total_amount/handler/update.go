package total_amount_handler

import (
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		old_category, _ := cmd.Flags().GetString("oldcategory")
		new_category, _ := cmd.Flags().GetString("newcategory")
		amount, _ := cmd.Flags().GetString("amount")
		label, _ := cmd.Flags().GetString("label")
		totalAmount := functions.StringToInt(amount)

		tv := structs.TotalAmountVariables{
			Included:    old_category,
			NewCategory: new_category,
			TotalAmount: totalAmount,
			Label:       label,
		}
		err := total_amount_db.UpdateTotalAmount(&tv)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UpdateCmd.Flags().StringP("oldcategory", "o", "", "Write the old category that you want to update")
	UpdateCmd.Flags().StringP("newcategory", "n", "", "Write the new category to update with")
	UpdateCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to update")
	UpdateCmd.Flags().StringP("label", "l", "", "Write the label that you want to update")
}
