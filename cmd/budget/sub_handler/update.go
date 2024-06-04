package budget_subhandler

import (
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the alert values for notification",
	Run: func(cmd *cobra.Command, args []string) {
		new_category, _ := cmd.Flags().GetString("new-category")
		old_category, _ := cmd.Flags().GetString("old-category")
		method, _ := cmd.Flags().GetString("method")
		frequency, _ := cmd.Flags().GetString("frequency")
		day, _ := cmd.Flags().GetString("day")
		weekday, _ := cmd.Flags().GetString("weekday")
		hour, _ := cmd.Flags().GetString("hour")
		minute, _ := cmd.Flags().GetString("minute")

		h := TakeHandler()
		dayInt := h.Deps.Common.StringToInt(day)
		hourInt := h.Deps.Common.StringToInt(hour)
		minuteInt := h.Deps.Common.StringToInt(minute)

		av := entities.AlertVariables{
			Category:    old_category,
			NewCategory: new_category,
			Method:      method,
			Frequency:   frequency,
			Days:        dayInt,
			Weekdays:    weekday,
			Hours:       hourInt,
			Minutes:     minuteInt,
		}

		err := h.Deps.AlertDB.UpdateAlert(&av)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UpdateCmd.Flags().StringP("new-category", "n", "", "Write the new category name to update it")
	UpdateCmd.Flags().StringP("old-category", "l", "", "Write the old category name to take its budget amount")
	UpdateCmd.Flags().StringP("frequency", "f", "", "Write the frequency of notifications (e.g., hourly, daily, weekly, monthly)")
	UpdateCmd.Flags().StringP("method", "t", "", "Write the preferred method of notification [email or CLI] message")
	UpdateCmd.Flags().StringP("day", "d", "", "Write the day to set the notification")
	UpdateCmd.Flags().StringP("weekday", "w", "", "Write the minute to set the notification")
	UpdateCmd.Flags().StringP("hour", "o", "", "Write the hour to set the notification")
	UpdateCmd.Flags().StringP("minute", "m", "", "Write the minute to set the notification")
}
