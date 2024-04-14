package total_amount_handler

import (
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// SetCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		include_category, _ := cmd.Flags().GetString("include")
		exclude_category, _ := cmd.Flags().GetString("exclude")
		label, _ := cmd.Flags().GetString("label")
		totalAmount := functions.StringToInt(amount)

		tv := structs.TotalAmountVariables{
			Amount:   totalAmount,
			Included: include_category,
			Excluded: exclude_category,
			Label:    label,
			Status:   "inactive",
		}

		err := total_amount_db.SetTotalAmount(&tv, "db/migrations/")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	SetCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to set")
	SetCmd.Flags().StringP("include", "i", "", "Specify a category to include in the total amount")
	SetCmd.Flags().StringP("exclude", "e", "", "Specify a category to exclude from the total amount")
	SetCmd.Flags().StringP("label", "l", "", "Provide a label for setting up your total amount")
}
