package budget_handler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the budget data in CSV",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("filepath")
		filename, _ := cmd.Flags().GetString("filename")

		myConnection := &db.MyConnection{}
		myBudget := &budget_db.MyBudgetDB{}
		myCommon := &conversion.MyCommon{}
		deps := interfaces.Dependencies{
			Connect:      myConnection,
			ManageBudget: myBudget,
			Common:       myCommon,
		}
		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myBudget.Handler = handle
		myCommon.Handler = handle

		err := handle.Deps.ManageBudget.GetBudgetData(filepath, filename)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	GetCmd.Flags().StringP("filepath", "p", "", "Give the file path to store the data")
	GetCmd.Flags().StringP("filename", "n", "", "Give the CSV file name to store the data")
}
