package budget_subhandler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/alert_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_alert "github.com/ibilalkayy/flow/usecases/app/alert"
	usecases_budget "github.com/ibilalkayy/flow/usecases/app/budget"
	"github.com/spf13/cobra"
)

// SetupCmd represents the setup command
var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set the alert notification",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		method, _ := cmd.Flags().GetString("method")
		frequency, _ := cmd.Flags().GetString("frequency")
		day, _ := cmd.Flags().GetString("day")
		weekday, _ := cmd.Flags().GetString("weekday")
		hour, _ := cmd.Flags().GetString("hour")
		minute, _ := cmd.Flags().GetString("minute")
		second, _ := cmd.Flags().GetString("second")

		myConnection := &db.MyConnection{}
		myAlert := &usecases_alert.MyAlert{}
		myAlertDB := &alert_db.MyAlertDB{}
		myBudget := &usecases_budget.MyBudget{}
		myCommon := &conversion.MyCommon{}
		deps := interfaces.Dependencies{
			Connect:      myConnection,
			ManageAlerts: myAlert,
			AlertDB:      myAlertDB,
			Budget:       myBudget,
			Common:       myCommon,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myAlert.Handler = handle
		myAlertDB.Handler = handle
		myCommon.Handler = handle
		myBudget.Handler = handle

		dayInt := handle.Deps.Common.StringToInt(day)
		hourInt := handle.Deps.Common.StringToInt(hour)
		minuteInt := handle.Deps.Common.StringToInt(minute)
		secondInt := handle.Deps.Common.StringToInt(second)

		av := entities.AlertVariables{
			Category:  category,
			Method:    method,
			Frequency: frequency,
			Days:      dayInt,
			Weekdays:  weekday,
			Hours:     hourInt,
			Minutes:   minuteInt,
			Seconds:   secondInt,
		}

		err := handle.Deps.ManageAlerts.AlertSetup(&av)
		if err != nil {
			log.Printf("Failed to setup the alert: %v", err)
		}
	},
}

func init() {
	SetupCmd.Flags().StringP("category", "c", "", "Write the category name to take its budget amount")
	SetupCmd.Flags().StringP("frequency", "f", "", "Write the frequency of notifications (e.g., hourly, daily, weekly, monthly)")
	SetupCmd.Flags().StringP("method", "t", "", "Write the preferred method of notification [email or CLI] message")
	SetupCmd.Flags().StringP("day", "d", "", "Write the day to set the notification")
	SetupCmd.Flags().StringP("weekday", "w", "", "Write a weekday to set the notification")
	SetupCmd.Flags().StringP("hour", "o", "", "Write the hour to set the notification")
	SetupCmd.Flags().StringP("minute", "m", "", "Write the minute to set the notification")
	SetupCmd.Flags().StringP("second", "s", "", "Write the second to set the notification")
}
