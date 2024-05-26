package spend_subhandler

import (
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the history data",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		myConnection := &db.MyConnection{}
		myHistory := &budget_db.MyBudgetDB{}
		deps := interfaces.Dependencies{
			Connect:      myConnection,
			ManageBudget: myHistory,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myHistory.Handler = handle

		err := handle.Deps.ManageBudget.RemoveHistory(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringP("category", "c", "", "Write the category to remove it's history")
}
