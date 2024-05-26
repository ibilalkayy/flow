package budget_subhandler

import (
	"log"

	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/alert_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the alert values for notification",
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
		myCommon := &conversion.MyCommon{}
		myAlertDB := &alert_db.MyAlertDB{}
		deps := interfaces.Dependencies{
			Connect: myConnection,
			AlertDB: myAlertDB,
			Common:  myCommon,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myAlertDB.Handler = handle
		myCommon.Handler = handle

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

		err := handle.Deps.AlertDB.UpdateAlert(&av)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UpdateCmd.Flags().StringP("category", "c", "", "Write the category name to take its budget amount")
	UpdateCmd.Flags().StringP("frequency", "f", "", "Write the frequency of notifications (e.g., hourly, daily, weekly, monthly)")
	UpdateCmd.Flags().StringP("method", "t", "", "Write the preferred method of notification [email or CLI] message")
	UpdateCmd.Flags().StringP("day", "d", "", "Write the day to set the notification")
	UpdateCmd.Flags().StringP("weekday", "w", "", "Write the minute to set the notification")
	UpdateCmd.Flags().StringP("hour", "o", "", "Write the hour to set the notification")
	UpdateCmd.Flags().StringP("minute", "m", "", "Write the minute to set the notification")
	UpdateCmd.Flags().StringP("second", "s", "", "Write the second to set the notification")
}
