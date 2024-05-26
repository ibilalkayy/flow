package budget_handler

import (
	"fmt"
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		oldCategory, _ := cmd.Flags().GetString("old-category")
		newCategory, _ := cmd.Flags().GetString("new-category")
		amount, _ := cmd.Flags().GetString("amount")

		myConnection := &db.MyConnection{}
		myCommon := &conversion.MyCommon{}
		myBudget := &budget_db.MyBudgetDB{}
		myTotalDB := &total_amount_db.MyTotalAmountDB{}
		deps := interfaces.Dependencies{
			Connect:             myConnection,
			TotalAmount:         myTotalDB,
			TotalAmountCategory: myTotalDB,
			ManageBudget:        myBudget,
			Common:              myCommon,
		}
		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myBudget.Handler = handle
		myTotalDB.Handler = handle
		myCommon.Handler = handle

		newAmount := handle.Deps.Common.StringToInt(amount)
		err := handle.Deps.ManageBudget.UpdateBudget(oldCategory, newCategory, newAmount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Your budget category is successfully updated!")
	},
}

func init() {
	UpdateCmd.Flags().StringP("old-category", "o", "", "Write the old category name to update")
	UpdateCmd.Flags().StringP("new-category", "n", "", "Write the new category name to allocate")
	UpdateCmd.Flags().StringP("amount", "a", "", "Write the new amount of the category to update")
}
