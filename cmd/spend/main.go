package spend

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/cmd"
	spend_handler "github.com/ibilalkayy/flow/cmd/spend/handler"
	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/blockchain"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/framework/email"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_spending "github.com/ibilalkayy/flow/usecases/app/spend"
	"github.com/ibilalkayy/flow/usecases/middleware"
	"github.com/spf13/cobra"
)

func TakeHandler() *handler.Handler {
	myConnection := &db.MyConnection{}
	myTotalDB := &total_amount_db.MyTotalAmountDB{}
	mySpendMoney := &usecases_spending.MySpending{}
	myBudgetAndHistory := &budget_db.MyBudgetDB{}
	myEmail := &email.MyEmail{}
	myEnv := &middleware.MyEnv{}
	myCommon := &conversion.MyCommon{}

	deps := interfaces.Dependencies{
		Connect:             myConnection,
		TotalAmount:         myTotalDB,
		TotalAmountCategory: myTotalDB,
		SpendAmount:         mySpendMoney,
		ManageBudget:        myBudgetAndHistory,
		HandleEmail:         myEmail,
		Env:                 myEnv,
		Common:              myCommon,
	}

	handle := handler.NewHandler(deps)
	myConnection.Handler = handle
	mySpendMoney.Handler = handle
	myTotalDB.Handler = handle
	myBudgetAndHistory.Handler = handle
	myEmail.Handler = handle
	myEnv.Handler = handle
	myCommon.Handler = handle

	return handle
}

// spendCmd represents the spend command
var SpendCmd = &cobra.Command{
	Use:   "spend",
	Short: "Spending money on various categories",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := blockchain.NewClient()
		if err != nil {
			log.Fatalf("Failed to create blockchain client: %v", err)
		}
		fmt.Println(client)

		categoryName, _ := cmd.Flags().GetString("category")
		spendingAmount, _ := cmd.Flags().GetString("amount")

		h := TakeHandler()
		spendingAmountInt := h.Deps.Common.StringToInt(spendingAmount)
		if len(categoryName) != 0 && spendingAmountInt != 0 {
			err := h.Deps.SpendAmount.SpendMoney(categoryName, spendingAmountInt)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(errors.New("select the command and flags"))
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(SpendCmd)
	// Subcommand
	SpendCmd.AddCommand(spend_handler.HistoryCmd)

	// Flags
	SpendCmd.Flags().StringP("category", "c", "", "Write the category name to spend the money on")
	SpendCmd.Flags().StringP("amount", "a", "", "Write the spending amount for a category")
}
