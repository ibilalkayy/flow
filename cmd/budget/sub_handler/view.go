package budget_subhandler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/alert_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// ViewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the alert notification values",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		myConnection := &db.MyConnection{}
		myAlertDB := &alert_db.MyAlertDB{}
		deps := interfaces.Dependencies{
			Connect: myConnection,
			AlertDB: myAlertDB,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myAlertDB.Handler = handle

		table, err := handle.Deps.AlertDB.ViewAlert(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}

func init() {
	ViewCmd.Flags().StringP("category", "c", "", "Write the category to see the alert notification values")
}
