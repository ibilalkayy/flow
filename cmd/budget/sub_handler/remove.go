package budget_subhandler

import (
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/alert_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the alert values",
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

		err := handle.Deps.AlertDB.RemoveAlert(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringP("category", "c", "", "Write the category name to remove its alert notification values")
}
