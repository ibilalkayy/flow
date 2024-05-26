package total_amount_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_total_amount "github.com/ibilalkayy/flow/usecases/app/total_amount"
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

		myConnection := &db.MyConnection{}
		myTotalDB := &total_amount_db.MyTotalAmountDB{}
		myTotal := &usecases_total_amount.MyTotalAmount{}
		myCommon := &conversion.MyCommon{}
		deps := interfaces.Dependencies{
			Connect:             myConnection,
			TotalAmount:         myTotalDB,
			TotalAmountCategory: myTotalDB,
			Total:               myTotal,
			Common:              myCommon,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myTotal.Handler = handle
		myTotalDB.Handler = handle
		myCommon.Handler = handle

		totalAmount := handle.Deps.Common.StringToInt(amount)
		err := handle.Deps.Total.SetTotalAmount(totalAmount, category, label)
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
