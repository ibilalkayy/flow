package budget_subhandler

import (
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/alert_db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/email"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_alert "github.com/ibilalkayy/flow/usecases/app/alert"
	usecases_spending "github.com/ibilalkayy/flow/usecases/app/spend"
	"github.com/ibilalkayy/flow/usecases/middleware"
	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for the alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		myConnection := &db.MyConnection{}
		myAlert := &usecases_alert.MyAlert{}
		myBudget := &budget_db.MyBudgetDB{}
		mySpending := &usecases_spending.MySpending{}
		myEmail := &email.MyEmail{}
		myAlertDB := &alert_db.MyAlertDB{}
		myEnv := &middleware.MyEnv{}

		deps := interfaces.Dependencies{
			Connect:      myConnection,
			ManageAlerts: myAlert,
			ManageBudget: myBudget,
			SpendAmount:  mySpending,
			HandleEmail:  myEmail,
			AlertDB:      myAlertDB,
			Env:          myEnv,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myAlert.Handler = handle
		myBudget.Handler = handle
		mySpending.Handler = handle
		myEmail.Handler = handle
		myAlertDB.Handler = handle
		myEnv.Handler = handle

		err := handle.Deps.ManageAlerts.CheckNotification(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	MsgCmd.Flags().StringP("category", "c", "", "Write the category to get the notification")
}
