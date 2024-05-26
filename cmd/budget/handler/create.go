package budget_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the budget of different categories",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetString("amount")

		myConnection := &db.MyConnection{}
		myBudget := &budget_db.MyBudgetDB{}
		myTotalDB := &total_amount_db.MyTotalAmountDB{}
		myCommon := &conversion.MyCommon{}
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

		amountInt := handle.Deps.Common.StringToInt(amount)
		bv := entities.BudgetVariables{Category: category, Amount: amountInt}

		err := handle.Deps.ManageBudget.CreateBudget(&bv)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	CreateCmd.Flags().StringP("category", "c", "", "Write the category like groceries, utilities, etc")
	CreateCmd.Flags().StringP("amount", "a", "", "Write the total amount for that category")
}
