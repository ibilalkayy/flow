package budget_handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// ViewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		myConnection := &db.MyConnection{}
		myBudget := &budget_db.MyBudgetDB{}
		deps := interfaces.Dependencies{
			Connect:      myConnection,
			ManageBudget: myBudget,
		}
		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myBudget.Handler = handle

		details, err := handle.Deps.ManageBudget.ViewBudget(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(details[0])
	},
}

func init() {
	ViewCmd.Flags().StringP("category", "c", "", "Write the category name to show the specific details")
}
