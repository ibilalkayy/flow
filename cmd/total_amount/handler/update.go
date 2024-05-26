package total_amount_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
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

		myConnection := &db.MyConnection{}
		myTotalDB := &total_amount_db.MyTotalAmountDB{}
		myCommon := &conversion.MyCommon{}

		deps := interfaces.Dependencies{
			Connect:             myConnection,
			TotalAmount:         myTotalDB,
			TotalAmountCategory: myTotalDB,
			Common:              myCommon,
		}
		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myTotalDB.Handler = handle
		myCommon.Handler = handle

		totalAmount := handle.Deps.Common.StringToInt(amount)
		tv := entities.TotalAmountVariables{
			Included:    old_category,
			NewCategory: new_category,
			TotalAmount: totalAmount,
			Label:       label,
		}

		err := handle.Deps.TotalAmount.UpdateTotalAmount(&tv)
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
