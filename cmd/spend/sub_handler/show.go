package spend_subhandler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// ShowCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the history data",
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

		table, err := handle.Deps.ManageBudget.ViewHistory(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}

func init() {
	ShowCmd.Flags().StringP("category", "c", "", "Write the category to show it's history")
}
